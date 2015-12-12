/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-12 16:55
 * description :
 * history :
 */

package partner

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jsix/gof"
	gfmt "github.com/jsix/gof/util/fmt"
	"github.com/jsix/gof/web"
	"go2o/src/core/domain/interface/member"
	"go2o/src/core/domain/interface/partner"
	"go2o/src/core/domain/interface/valueobject"
	"go2o/src/core/infrastructure/format"
	"go2o/src/core/service/dps"
	"go2o/src/core/variable"
	"html/template"
	"strconv"
	"time"
	"go2o/src/x/echox"
	"net/http"
)

type memberC struct {
}

func (this *memberC) LevelList(ctx *echox.Context) error {
	d := echox.NewRenderData()
	return ctx.RenderOK("member/level_list.html", d)
}

//修改门店信息
func (this *memberC) EditMLevel(ctx *echox.Context) error {
	partnerId := getPartnerId(ctx)
	id, _ := strconv.Atoi(ctx.Query("id"))
	entity := dps.PartnerService.GetMemberLevelById(partnerId, id)
	js, _ := json.Marshal(entity)
	d := echox.NewRenderData()
	d.Map["entity"] = template.JS(js)
	return ctx.RenderOK("member/edit_level.html", d)
}

func (this *memberC) CreateMLevel(ctx *echox.Context) error {
	d := echox.NewRenderData()
	d.Map["entity"] = template.JS("{}")
	return ctx.RenderOK("member/create_level.html", d)
}

func (this *memberC) SaveMLevel_post(ctx *echox.Context) error {
	partnerId := getPartnerId(ctx)
	r := ctx.Request
	var result gof.Message
	r.ParseForm()

	e := valueobject.MemberLevel{}
	web.ParseFormToEntity(r.Form, &e)
	e.PartnerId = getPartnerId(ctx)

	id, err := dps.PartnerService.SaveMemberLevel(partnerId, &e)

	if err != nil {
		result.Message = err.Error()
	} else {
		result.Result = true
		result.Data = id
	}
	return ctx.JSON(http.StatusOK, result)
}

func (this *memberC) DelMLevel(ctx *echox.Context) error {
	r := ctx.Request
	var result gof.Message
	r.ParseForm()
	partnerId := getPartnerId(ctx)
	id, err := strconv.Atoi(r.FormValue("id"))
	if err == nil {
		err = dps.PartnerService.DelMemberLevel(partnerId, id)
	}

	if err != nil {
		result.Message = err.Error()
	} else {
		result.Result = true
	}
	return ctx.JSON(http.StatusOK, result)
}

// 会员列表
func (this *memberC) List(ctx *echox.Context) error {
	levelDr := getLevelDropDownList(getPartnerId(ctx))
	d := echox.NewRenderData()
	d.Map["levelDr"] = template.HTML(levelDr)
	return ctx.RenderOK("member/member_list.html", d)
}

// 锁定会员
func (this *memberC) Lock_member_post(ctx *echox.Context) error {
	ctx.Request.ParseForm()
	id, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	partnerId := getPartnerId(ctx)
	var result gof.Message
	if _, err := dps.MemberService.LockMember(partnerId, id); err != nil {
		result.Message = err.Error()
	} else {
		result.Result = true
	}
	return ctx.JSON(http.StatusOK, result)
}

func (this *memberC) Member_details(ctx *echox.Context) error {
	memberId, _ := strconv.Atoi(ctx.Request.URL.Query().Get("member_id"))

	d := echox.NewRenderData()
	d.Map["memberId"] = memberId
	return ctx.RenderOK("member/member_details.html", d)
}

// 会员基本信息
func (this *memberC) Member_basic(ctx *echox.Context) error {
	memberId, _ := strconv.Atoi(ctx.Request.URL.Query().Get("member_id"))
	m := dps.MemberService.GetMember(memberId)
	if m == nil {
		return ctx.String(http.StatusOK, "no such member")
	}
	lv := dps.PartnerService.GetLevel(getPartnerId(ctx), m.Level)
	d := echox.NewRenderData()
	d.Map = map[string]interface{}{
		"m":  m,
		"lv": lv,
		"sexName": gfmt.BoolString(m.Sex == 1, "先生",
			gfmt.BoolString(m.Sex == 2, "女士", "-")),
		"lastLoginTime": format.HanUnixDateTime(m.LastLoginTime),
		"regTime":       format.HanUnixDateTime(m.RegTime),
	}

	return ctx.RenderOK("member/basic_info.html", d)
}

// 会员账户信息
func (this *memberC) Member_account(ctx *echox.Context) error {
	memberId, _ := strconv.Atoi(ctx.Request.URL.Query().Get("member_id"))
	acc := dps.MemberService.GetAccount(memberId)
	if acc == nil {
		return ctx.String(http.StatusOK,"no such account")
	}

		d := echox.NewRenderData()
		d.Map = map[string]interface{}{
				"acc": acc,
				"balanceAccountAlias": variable.AliasBalanceAccount,
				"presentAccountAlias": variable.AliasPresentAccount,
				"flowAccountAlias":    variable.AliasFlowAccount,
				"growAccountAlias":    variable.AliasGrowAccount,
				"integralAlias":       variable.AliasIntegral,
				"updateTime":          format.HanUnixDateTime(acc.UpdateTime),
			}
		return ctx.Render(http.StatusOK,"member/account_info.html",d)


}

// 会员收款银行信息
func (this *memberC) Member_curr_bank(ctx *echox.Context) error {
	memberId, _ := strconv.Atoi(ctx.Request.URL.Query().Get("member_id"))
	e := dps.MemberService.GetBank(memberId)
	if e != nil && len(e.Account) > 0 && len(e.AccountName) > 0 &&
	len(e.Name) > 0 && len(e.Network) > 0 {
		d := echox.NewRenderData()
		d.Map["bank"] = e
		return ctx.RenderOK("member/member_curr_bank.html", d)
	}
	return ctx.String(http.StatusOK, "<span class=\"red\">尚未完善</span>")
}

func (this *memberC) Reset_pwd_post(ctx *echox.Context) error {
	var result gof.Message
	ctx.Request.ParseForm()
	memberId, _ := strconv.Atoi(ctx.Request.FormValue("member_id"))
	rl := dps.MemberService.GetRelation(memberId)
	partnerId := getPartnerId(ctx)
	if rl == nil || rl.RegisterPartnerId != partnerId {
		result.Message = "无权进行当前操作"
	} else {
		newPwd := dps.MemberService.ResetPassword(memberId)
		result.Result = true
		result.Message = fmt.Sprintf("重置成功,新密码为: %s", newPwd)
	}
	return ctx.JSON(http.StatusOK, result)
}

// 客服充值
func (this *memberC) Charge(ctx *echox.Context) error {
	memberId, _ := strconv.Atoi(ctx.Request.URL.Query().Get("member_id"))
	mem := dps.MemberService.GetMemberSummary(memberId)
	if mem == nil {
		return ctx.String(http.StatusOK, "no such member")
	}
	d := echox.NewRenderData()
	d.Map["m"] = mem
	return ctx.RenderOK("member/charge.html", d)

}

func (this *memberC) Charge_post(ctx *echox.Context) error {
	var msg gof.Message
	var err error
	ctx.Request.ParseForm()
	partnerId := getPartnerId(ctx)
	memberId, _ := strconv.Atoi(ctx.Request.FormValue("MemberId"))
	amount, _ := strconv.ParseFloat(ctx.Request.FormValue("Amount"), 32)
	if amount < 0 {
		msg.Message = "error amount"
	} else {
		rel := dps.MemberService.GetRelation(memberId)

		if rel.RegisterPartnerId != getPartnerId(ctx) {
			err = partner.ErrPartnerNotMatch
		} else {
			title := fmt.Sprintf("客服充值%f", amount)
			err = dps.MemberService.Charge(partnerId, memberId, member.TypeBalanceServiceCharge, title, "", float32(amount))
		}
		if err != nil {
			msg.Message = err.Error()
		} else {
			msg.Result = true
		}
	}
	return ctx.JSON(http.StatusOK, msg)
}

// 提现列表
func (this *memberC) ApplyRequestList(ctx *echox.Context) error {
	levelDr := getLevelDropDownList(getPartnerId(ctx))

	d := echox.NewRenderData()
	d.Map["levelDr"] = template.HTML(levelDr)
	d.Map["kind"] =  member.KindBalanceApplyCash
	return ctx.RenderOK("member/apply_request_list.html", d)
}

// 审核提现请求
func (this *memberC) Pass_apply_req_post(ctx *echox.Context) error {
	var msg gof.Message
	ctx.Request.ParseForm()
	partnerId := getPartnerId(ctx)
	passed := ctx.Request.FormValue("pass") == "1"
	memberId, _ := strconv.Atoi(ctx.Request.FormValue("member_id"))
	id, _ := strconv.Atoi(ctx.Request.FormValue("id"))

	err := dps.MemberService.ConfirmApplyCash(partnerId, memberId, id, passed, "")

	if err != nil {
		msg.Message = err.Error()
	} else {
		msg.Result = true
	}
	return ctx.JSON(http.StatusOK, msg)
}

// 退回提现请求
func (this *memberC) Back_apply_req(ctx *echox.Context) error {
	form := ctx.Request.URL.Query()
	memberId, _ := strconv.Atoi(form.Get("member_id"))
	id, _ := strconv.Atoi(form.Get("id"))

	info := dps.MemberService.GetBalanceInfoById(memberId, id)

	if info == nil {
		return ctx.String(http.StatusOK, "no such request")
	}

	d := echox.NewRenderData()
	d.Map["info"] = info
	d.Map["applyTime"] = time.Unix(info.CreateTime, 0).Format("2006-01-02 15:04:05")
	return ctx.RenderOK("member/back_apply_req.html", d)
}


func (this *memberC) Back_apply_req_post(ctx *echox.Context) error {
	var msg gof.Message
	ctx.Request.ParseForm()
	form := ctx.Request.Form
	partnerId := getPartnerId(ctx)
	memberId, _ := strconv.Atoi(form.Get("MemberId"))
	id, _ := strconv.Atoi(form.Get("Id"))

	err := dps.MemberService.ConfirmApplyCash(partnerId, memberId, id, false, "")
	if err != nil {
		msg.Message = err.Error()
	} else {
		msg.Result = true
	}
	return ctx.JSON(http.StatusOK, msg)
}

// 提现打款
func (this *memberC) Handle_apply_req(ctx *echox.Context) error {
	form := ctx.Request.URL.Query()
	memberId, _ := strconv.Atoi(form.Get("member_id"))
	id, _ := strconv.Atoi(form.Get("id"))

	info := dps.MemberService.GetBalanceInfoById(memberId, id)

	if info == nil {
		return ctx.String(http.StatusOK, "no such info")
	}

	d := echox.NewRenderData()
	bank := dps.MemberService.GetBank(memberId)
	d.Map = map[string]interface{}{
		"info":      info,
		"bank":      bank,
		"applyTime": time.Unix(info.CreateTime, 0).Format("2006-01-02 15:04:05"),
	}
	return ctx.RenderOK("member/handle_apply_req.html", d)
}

func (this *memberC) Handle_apply_req_post(ctx *echox.Context) error {
	var msg gof.Message
	var err error
	ctx.Request.ParseForm()
	form := ctx.Request.Form
	partnerId := getPartnerId(ctx)
	memberId, _ := strconv.Atoi(form.Get("MemberId"))
	id, _ := strconv.Atoi(form.Get("Id"))
	agree := form.Get("Agree") == "on"
	tradeNo := form.Get("TradeNo")

	if !agree {
		err = errors.New("请同意已知晓并打款选项")
	} else {
		err = dps.MemberService.FinishApplyCash(partnerId, memberId, id, tradeNo)
	}
	if err != nil {
		msg.Message = err.Error()
	} else {
		msg.Result = true
	}
	return ctx.JSON(http.StatusOK, msg)
}

// 团队排名列表
func (this *memberC) Team_rank(ctx *echox.Context) error {

	levelDr := getLevelDropDownList(getPartnerId(ctx))
	d := echox.NewRenderData()
	d.Map["levelDr"] = template.HTML(levelDr)
	return ctx.RenderOK("member/team_rank.html", d)
}
