package thirdPayment

import (
	"context"
	"net/http"

	"aiops/app/payment/cmd/api/internal/svc"
	"aiops/app/payment/cmd/api/internal/types"
	"aiops/app/payment/cmd/rpc/payment"
	"aiops/app/payment/model"
	"aiops/common/xerr"

	"github.com/pkg/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrWxPayCallbackError = xerr.NewErrMsg("wechat pay callback fail")

type ThirdPaymentcallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type ThirdPaymentWxPayCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThirdPaymentWxPayCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) ThirdPaymentWxPayCallbackLogic {
	return ThirdPaymentWxPayCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPaymentWxPayCallbackLogic) ThirdPaymentWxPayCallback(rw http.ResponseWriter, req *http.Request) (*types.ThirdPaymentWxPayCallbackResp, error) {

	//Retrieve the local merchant certificate private key.
	_, err := svc.NewWxPayClientV3(l.svcCtx.Config)
	if err != nil {
		return nil, err
	}

	// Get platform certificate accessor
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(l.svcCtx.Config.WxPayConf.MchId)
	handler := notify.NewNotifyHandler(l.svcCtx.Config.WxPayConf.APIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))
	//Verifying signatures, parsing data
	transaction := new(payments.Transaction)
	_, err = handler.ParseNotifyRequest(context.Background(), req, transaction)
	if err != nil {
		return nil, errors.Wrapf(ErrWxPayCallbackError, "Failed to parse data ,err:%v", err)
	}

	returnCode := "SUCCESS"
	err = l.verifyAndUpdateState(transaction)
	if err != nil {
		returnCode = "FAIL"
	}

	return &types.ThirdPaymentWxPayCallbackResp{
		ReturnCode: returnCode,
	}, err

}

//Verify and update relevant flow data
func (l *ThirdPaymentWxPayCallbackLogic) verifyAndUpdateState(notifyTrasaction *payments.Transaction) error {

	paymentResp, err := l.svcCtx.PaymentRpc.GetPaymentBySn(l.ctx, &payment.GetPaymentBySnReq{
		Sn: *notifyTrasaction.OutTradeNo,
	})
	if err != nil || paymentResp.PaymentDetail.Id == 0 {
		return errors.Wrapf(ErrWxPayCallbackError, "Failed to get payment flow record err:%v ,notifyTrasaction:%+v ", err, notifyTrasaction)
	}

	//????????????
	notifyPayTotal := *notifyTrasaction.Amount.PayerTotal
	if paymentResp.PaymentDetail.PayTotal != notifyPayTotal {
		return errors.Wrapf(ErrWxPayCallbackError, "Order amount exception  notifyPayTotal:%v , notifyTrasaction:%v ", notifyPayTotal, notifyTrasaction)
	}

	// Judgment status
	payStatus := l.getPayStatusByWXPayTradeState(*notifyTrasaction.TradeState)
	if payStatus == model.ThirdPaymentPayTradeStateSuccess {
		//Payment Notification.

		if paymentResp.PaymentDetail.PayStatus != model.ThirdPaymentPayTradeStateWait {
			return nil
		}

		// Update the flow status.
		if _, err = l.svcCtx.PaymentRpc.UpdateTradeState(l.ctx, &payment.UpdateTradeStateReq{
			Sn:             *notifyTrasaction.OutTradeNo,
			TradeState:     *notifyTrasaction.TradeState,
			TransactionId:  *notifyTrasaction.TransactionId,
			TradeType:      *notifyTrasaction.TradeType,
			TradeStateDesc: *notifyTrasaction.TradeStateDesc,
			PayStatus:      l.getPayStatusByWXPayTradeState(*notifyTrasaction.TradeState),
		}); err != nil {
			return errors.Wrapf(ErrWxPayCallbackError, "????????????????????????  err:%v , notifyTrasaction:%v ", err, notifyTrasaction)
		}

	} else if payStatus == model.ThirdPaymentPayTradeStateWait {
		//Refund notification @todo to be done later, not needed at this time
	}

	return nil

}

const (
	SUCCESS    = "SUCCESS"    //????????????
	REFUND     = "REFUND"     //????????????
	NOTPAY     = "NOTPAY"     //?????????
	CLOSED     = "CLOSED"     //?????????
	REVOKED    = "REVOKED"    //??????????????????????????????
	USERPAYING = "USERPAYING" //????????????????????????????????????
	PAYERROR   = "PAYERROR"   //????????????(????????????????????????????????????)
)

func (l *ThirdPaymentWxPayCallbackLogic) getPayStatusByWXPayTradeState(wxPayTradeState string) int64 {

	switch wxPayTradeState {
	case SUCCESS: //????????????
		return model.ThirdPaymentPayTradeStateSuccess
	case USERPAYING: //?????????
		return model.ThirdPaymentPayTradeStateWait
	case REFUND: //?????????
		return model.ThirdPaymentPayTradeStateWait
	default:
		return model.ThirdPaymentPayTradeStateFAIL
	}
}
