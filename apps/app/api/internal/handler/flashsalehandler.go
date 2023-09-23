package handler

import (
	"net/http"

	"github.com/wangzhou-ccc/mygozero/apps/app/api/internal/logic"
	"github.com/wangzhou-ccc/mygozero/apps/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FlashSaleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFlashSaleLogic(r.Context(), svcCtx)
		resp, err := l.FlashSale()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
