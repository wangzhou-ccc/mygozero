package user

import (
	"net/http"

	"github.com/wangzhou-ccc/mygozero/apps/app/api/internal/logic/user"
	"github.com/wangzhou-ccc/mygozero/apps/app/api/internal/svc"
	"github.com/wangzhou-ccc/mygozero/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserReceiveAddressListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReceiveAddressListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserReceiveAddressListLogic(r.Context(), svcCtx)
		resp, err := l.UserReceiveAddressList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
