package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

	chrome "github.com/mkenney/go-chrome/tot"
	"github.com/mkenney/go-chrome/tot/cdtp/emulation"
	"github.com/mkenney/go-chrome/tot/cdtp/page"
	"github.com/mkenney/go-chrome/tot/socket"
	logfmt "github.com/mkenney/go-log-fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
	level, _ := log.ParseLevel("debug")
	log.SetLevel(level)
	log.SetFormatter(&logfmt.TextFormat{})
}

func main() {
	var err error

	// Define a chrome instance with remote debugging enabled.
	browser := chrome.New(
		// See https://developers.google.com/web/updates/2017/04/headless-chrome#cli
		// for details about startup flags
		&chrome.Flags{
			"addr":               "localhost",
			"disable-extensions": nil,
			"disable-gpu":        nil,
			"headless":           nil,
			"hide-scrollbars":    nil,
			"no-first-run":       nil,
			"no-sandbox":         nil,
			"port":               9222,
			"remote-debugging-address": "0.0.0.0",
			"remote-debugging-port":    9222,
		},
		"/usr/bin/google-chrome", // Path to Chromeium binary
		"/tmp",      // Set the Chromium working directory
		"/dev/null", // Ignore internal Chromium output, set to empty string for os.Stdout
		"/dev/null", // Ignore internal Chromium errors, set to empty string for os.Stderr
	)

	// Start the chrome process.
	if err := browser.Launch(); nil != err {
		panic(err)
	}

	// Open a tab and navigate to the URL you want to screenshot.
	tab, err := browser.NewTab("")
	if nil != err {
		panic(err)
	}

	// Enable Page events for this tab.
	if enableResult := <-tab.Page().Enable(); nil != enableResult.Err {
		panic(enableResult.Err)
	}

	var frameID page.FrameID
	ftResult := <-tab.Page().GetFrameTree()
	tmp, _ := json.MarshalIndent(ftResult, "", "    ")
	log.Debugf("Frame Tree: %s", string(tmp))
	log.Debugf("Err: %+v", ftResult.Err)
	setContentResult := <-tab.Page().SetDocumentContent(&page.SetDocumentContentParams{
		FrameID: frameID,
		HTML: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
		<html xmlns="http://www.w3.org/1999/xhtml">
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<meta name="format-detection" content="telephone=no" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		<title>Order Confirmation</title>
		<style type="text/css">
			/* Forces Hotmail to display emails at full width */
			.ReadMsgBody {
				width: 100%;
			}
			/*Hotmail table centering fix*/
			.ExternalClass {
				width: 100%;
			}
			/* Forces Hotmail/Outlook to display normal line spacing*/
			.ExternalClass * {
				line-height: 100%;
			}
			/*Yahoo paragraph fix*/
			p {
				margin: 1em 0;
			}
			/*This resolves the Outlook 07, 10, and Gmail td padding issue fix*/
			table td {
				border-collapse: collapse;
			}
			/*viewport width scaling for all devices*/
			 @-ms-viewport {
			 width: device-width;
			}
			/*This resolves the issue when iphone puts links on dates, etc.*/
			.appleLinks {
				color: inherit;
				text-decoration: none;
			}
			.appleLinks a {
				color: inherit;
				text-decoration: none;
			}

			body[yahoo] .autoht, .header-myaccount, .header-customerservice {
				height: inherit !important;
			}

			/*Mobile - Media Queries Starts - 320X480 Screen*/
			@media screen and (max-width: 480px) {
				/* shared content area CSS */
			body[yahoo] .header-mobile-hidden {
					display: none !important;
				}
				body[yahoo] .header-logoimg {
					width: 121px !important;
					height: inherit !important;
				}
				body[yahoo] .header-logo {
					padding-right: 10px !important;
				}
				body[yahoo] .header-bopusimg {
					width: 100% !important;
					height: auto !important;
					max-width: 239px !important;
				}
				body[yahoo] .header-mobilenav {
					width: 156px !important;
				}
				body[yahoo] .header-myaccount {
					width: 60px !important;
					height: inherit !important;
				}
				body[yahoo] .header-customerservice {
					width: 92px !important;
					height: inherit !important;
				}
				body[yahoo] .header-menuborder {
					height: 10px !important;
				}
				body[yahoo] .header-menufont, .header-menufont a, .header-menufont a span {
					font-size: 11px !important;
					color: #000001 !important;
				}
				body[yahoo] .header-menufontlast, .header-menufontlast a, .header-menufontlast a span {
					font-size: 11px !important;
					color: #fd232b !important;
				}
				body[yahoo] .header-menupad {
					padding: 0 8px 0 8px !important;
				}
				body[yahoo] .header-menupad-first {
					word-wrap: normal !important;
					padding: 0 8px 0 0 !important;
				}
				body[yahoo] .header-menupad-last {
					word-wrap: normal !important;
					padding: 0 0 0 8px !important;
				}
			body[yahoo] .footer-drop {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
				}
				body[yahoo] .footer-fullwidth {
					width: 100% !important;
				}
				body[yahoo] .footer-noborder {
					border: 0 !important;
				}
				body[yahoo] .footer-toppad {
					padding-top: 20px !important;
				}
				body[yahoo] .footer-width-auto {
					width: auto !important;
				}
				body[yahoo] .footer-width90 {
					width: 90% !important;
				}
				body[yahoo] .footer-nopadtp {
					padding-top: 0 !important;
				}
			body[yahoo] .shipping-zeropadrt {
					padding-right: 0px !important;
				}
				body[yahoo] .shipping-drop {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
				}
				body[yahoo] .shipping-item-fs15 {
					font-size: 15px !important;
				}
				body[yahoo] .shipping-fnt12 {
					font-size: 12px !important;
				}
				body[yahoo] .shipping-item-mobpadhdr {
					padding-top: 15px !important;
				}
				body[yahoo] .shipping-mobpad {
					padding: 20px 15px 25px 15px !important;
				}
				body[yahoo] .shipping-bordermob {
					border-bottom: 1px solid #cccccc !important;
					border-top: 1px solid #cccccc !important;
				}
				body[yahoo] .shipping-pickup-mobpickpad {
					padding-bottom: 12px !important;
				}
				body[yahoo] .shipping-widthaddr {
					max-width: 150px !important;
				}
				body[yahoo] .shipping-col46 {
					color: #464646 !important;
				}
				body[yahoo] .shipping-top {
					padding-top: 30px !important;
				}
				body[yahoo] .shipping-item-gwp {
					padding: 14px 0 !important;
				}
				body[yahoo] .shipping-lhmob {
					line-height: 13px !important;
				}
				body[yahoo] .shipping-width-100 {
					width: 100% !important;
				}
				body[yahoo] .shipping-nopadlft {
					padding-left: 0px !important;
				}
				body[yahoo] .shipping-fnt11 {
					font-size: 11px !important;
				}
				body[yahoo] .shipping-txtalignleft {
					text-align: left !important;
					padding-left: 3px !important;
				}
				body[yahoo] .shipping-width130left {
					width: 130px !important;
					height: auto; !important
					padding-right: 28px !important;
				}
				body[yahoo] .shipping-mobpad-sides {
					padding: 0 15px 0 15px !important;
				}
				body[yahoo] .shipping-width130 {
					max-width: 130px !important;
				}
				body[yahoo] .shipping-width-ed {
					max-width: 140px !important;
				}
				body[yahoo] .shipping-padtpbt {
				}
				body[yahoo] .shipping-btnmaxwidth {
				}
				body[yahoo] .shipping-bold {
				}
			body[yahoo] .product-fnt12 {
					font-size: 12px !important;
				}
				body[yahoo] .product-item-fs15 {
					font-size: 15px !important;
				}
				body[yahoo] .product-display-inline {
					display: inline !important;
					float: left !important;
				}
				body[yahoo] .product-display-inline-right {
					float: right !important;
					display: inline !important;
				}
				body[yahoo] .product-width130right {
					width: 130px !important;
					height: 130px; !important
				}
				body[yahoo] .product-width130left {
					width: 130px !important;
					height: 130px; !important
					padding-right: 28px !important;
				}
				body[yahoo] .product-width130 {
					max-width: 130px !important;
				}
				body[yahoo] .product-zeropadrt {
					padding-right: 0px !important;
				}
				body[yahoo] .product-zeropad {
					padding: 0px !important;
				}
				body[yahoo] .product-drop {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
				}
				body[yahoo] .product-mid-contpad {
					padding-top: 15px;
					padding-bottom: 22px;
					font-size: 12px !important;
					font-weight: bold !important;
				}
				body[yahoo] .product-mid-contpad3 {
					padding-bottom: 7px;
					font-size: 12px !important;
					font-weight: bold !important;
				}
				body[yahoo] .product-lhmob {
					line-height: 13px !important;
				}
				body[yahoo] .product-item-gwp {
					padding: 14px 0 !important;
				}
				body[yahoo] .product-top {
					padding-top: 30px !important;
				}
				body[yahoo] .product-yppad {
					padding-top: 14px !important;
					text-align: left !important;
					font-size: 14px !important;
				}
				body[yahoo] .product-width-100 {
					width: 100% !important;
				}
				body[yahoo] .product-nopadlft {
					padding-left: 0px !important;
				}
				body[yahoo] .product-fnt11 {
					font-size: 11px !important;
				}
				body[yahoo] .product-txtalignleft {
					text-align: left !important;
					padding-left: 3px !important;
				}
				body[yahoo] .product-checkmark {
					padding-top: 0 !important;
				}
			body[yahoo] .help-drop {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
				}
				body[yahoo] .help-nopadtp {
					padding-top: 0 !important;
				}
				body[yahoo] .help-nhpad {
					padding-top: 5px !important;
					padding-bottom: 15px !important;
				}
				body[yahoo] .help-width288 {
					width: 288px !important;
				}
				body[yahoo] .help-padgpick {
					padding-top: 35px !important;
				}
				body[yahoo] .help-fnt19 {
					font-size: 19px !important;
				}
				body[yahoo] .help-fnt16 {
					font-size: 16px !important;
				}
				body[yahoo] .help-fnt13 {
					font-size: 13px !important;
				}
			body[yahoo] .helpmkt-drop {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
				}
				body[yahoo] .helpmkt-nopadtp {
					padding-top: 0 !important;
				}
				body[yahoo] .helpmkt-nhpad {
					padding-top: 5px !important;
					padding-bottom: 15px !important;
				}
				body[yahoo] .helpmkt-width288 {
					width: 288px !important;
				}
				body[yahoo] .helpmkt-padgpick {
					padding-top: 35px !important;
				}
				body[yahoo] .helpmkt-fnt19 {
					font-size: 19px !important;
				}
				body[yahoo] .helpmkt-fnt16 {
					font-size: 16px !important;
				}
				body[yahoo] .helpmkt-fnt13 {
					font-size: 13px !important;
				}
			body[yahoo] .summary-drop {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
				}
				body[yahoo] .summary-nopadtp {
					padding-top: 0 !important;
				}
				body[yahoo] .summary-item-fs15 {
					font-size: 15px !important;
				}
				body[yahoo] .summary-colorblk {
					color: #000001 !important;
				}
				body[yahoo] .summary-padrtzero {
					padding-right: 0px !important;
				}
				body[yahoo] .summary-fnt13 {
					font-size: 13px !important;
				}
				body[yahoo] .summary-item-mobpadhdr {
					padding-top: 15px !important;
					padding-left: 0 !important;
				}
				body[yahoo] .summary-item-mobpadbottom {
					padding-bottom: 15px !important;
				}
				body[yahoo] .summary-ordersummary-padbtos {
					padding-bottom: 12px !important;
				}
				body[yahoo] .summary-earning-padbtye {
					padding-bottom: 12px !important;
					padding-top: 40px !important;
				}
				body[yahoo] .summary-yrstp {
					padding-top: 20px !important;
				}
				body[yahoo] .summary-yrs {
					font-size: 16px !important;
				}
			body[yahoo] .recs-mobile-hidden {
				display:none !important;
				padding: 0 !important;
				}
				body[yahoo] .recs-mobile-show {
					display: block !important;
					margin: 0 !important;
					padding: 0 !important;
					overflow: visible !important;
					width: auto !important;
					max-height: inherit !important;
				}
			  body[yahoo] img.recs-header-img {
				width:290px !important;
				display:block !important;
				height:auto !important;
				float:none !important;
			  }
				body[yahoo] .recs-nopadtp {
					padding-top: 0 !important;
				}
				body[yahoo] .recs-fnt19 {
					font-size: 19px !important;
				}
				body[yahoo] .recs-alg-center {
					text-align: center !important;
				}
				body[yahoo] .recs-drop {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
				}
			body[yahoo] .ads-mobile-hidden {
					display:none !important;
					padding: 0 !important;
				}
				body[yahoo] .ads-mobile-show {
					display: block !important;
					margin: 0 !important;
					padding: 0 !important;
					overflow: visible !important;
					width: auto !important;
					max-height: inherit !important;
				}
				body[yahoo] img.ads-header-img {
					width:290px !important;
					display:block !important;
					height:auto !important;
					float:none !important;
				}
				body[yahoo] .ads-nopadtp {
					padding-top: 0 !important;
				}
				body[yahoo] .ads-fnt19 {
					font-size: 19px !important;
				}
				body[yahoo] .ads-alg-center {
					text-align: center !important;
				}
				body[yahoo] .ads-drop {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
					padding: 20px 0px 0px 0px !important;
				}
				body[yahoo] .ads-wide {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
					padding: 20px 0px 0px 0px !important;
				}

				/*Do not show in mobiles*/
				body[yahoo] .mobile-hidden {
					display: none !important;
				}
				/*Shows in mobiles*/
				body[yahoo] .mobile-show {
					display: block !important;
					margin: 0 !important;
					padding: 0 !important;
					overflow: visible !important;
					width: auto !important;
					max-height: inherit !important;
				}
				/*Table,Td resize*/
				table[class=resize_table_to_320], td[class=resize_table_to_320] {
					width: 100% !important;
					height: auto;
					margin: 0 auto;
				}
				body[yahoo] .padbot25 {
					padding-bottom: 25px;
				}
				body[yahoo] .drop {
					width: 100% !important;
					float: left !important;
					display: block !important;
					text-align: left !important;
				}
				body[yahoo] .photo {
					width: 100% !important;
					height: auto !important;
				}
				body[yahoo] .spacer15 {
					width: 1px !important;
					height: 15px !important;
				}
				body[yahoo] .spacer30 {
					width: 1px !important;
					height: 30px !important;
				}
				body[yahoo] .fnt12 {
					font-size: 12px !important;
				}
				body[yahoo] .prod-mobpad {
					padding: 0 15px 25px 15px !important;
				}
				body[yahoo] .nopadtpnbt {
					padding-top: 0px !important;
					padding-bottom: 0 !important;
				}
				body[yahoo] .colorblk {
					color: #000001 !important;
				}
				body[yahoo] .padrt15 {
					padding-right: 15px !important;
				}
				body[yahoo] .padrt13 {
					padding-right: 14px !important;
				}
				body[yahoo] .alg-center {
					text-align: center !important;
				}
				body[yahoo] .moborderpad {
					padding-top: 20px !important;
				}
				body[yahoo] .mobbotpadbtn {
					padding-bottom: 20px;
				}
				body[yahoo] .fnt20 {
					font-size: 20px !important;
					line-height: 20px !important;
				}
				body[yahoo] .nopadtop {
					padding-top: 0px !important;
				}
				body[yahoo] .nopadlr {
					padding-left: 0px !important;
					padding-right: 0px !important;
				}
				body[yahoo] .width-ed {
					max-width: 140px !important;
				}
				body[yahoo] .width140 {
					max-width: 140px !important;
				}
				body[yahoo] .padtopzero {
					padding-top: 0 !important;
				}
				body[yahoo] .maxwidth-content {
					max-width: 350px !important;
				}
				body[yahoo] .incentive {
					padding: 10px 20px !important;
				}
			}
		</style>
		</head>
		<body class="appleLinks" bgcolor="#FFFFFF" style="margin: 0px; padding:0px; -webkit-text-size-adjust:none; -ms-text-size-adjust: none;line-height:16px; mso-line-height-rule:exactly;" yahoo="fix">
		<style type="text/css">
		div.preheader
		{ display: none !important; }
		</style>
		<div class="preheader" style="font-size: 1px; display: none !important;">Don't leave yet ░ your order ░░░░ be ready for pickup ░░░░.</div>

		<!-- ET Tracking Code-->
		<img src="http://click.t.kohls.com/open.aspx?ffcb10-fed110717460057e-fe6115797566017f7c16-fe8a1272716d007476-ff951579-fe891d757463067a72-febd15777d620378" width="1" height="1">
		<!-- Table used to center email Starts-->
		<table border="0" cellpadding="0" cellspacing="0" width="100%">
		  <tr>
			<td align="left" valign="top" style="padding: 15px 0px 20px 0px; font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #888888;" bgcolor="#ffffff" class="nopadtop">
			  <div align="center">
				<!-- Table used to set width of email -->
				<table class="resize_table_to_320" border="0" cellpadding="0" cellspacing="0" width="640" bgcolor="#ffffff">

				  <!--Header starts-->
				  <!--VAWP starts-->
				  <tr>
					<td align="right" valign="top" style="padding: 10px 10px 10px 0;">
					  <a href="http://view.t.kohls.com/?qs=8a661ca3adf24beb4bf61353a0e9bbc9b07008783746982f2e036b450a4cc855cd9b184ad0c55e8545a4346ddee80baddf40d66078361873e8c0852a67880337eadd1df5c0721c8e8accda5f3533c092" target="_blank" style="color: #00add8; text-decoration: none;" ><span style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545; text-decoration: underline;">View this email in a browser</span></a>
					</td>
				  </tr>
				  <!--VAWP ends-->

				  <!--Header starts-->
				  <tr>
					<td align="center" valign="top" bgcolor="#454545" style="padding-top:10px; padding-bottom:8px;">
					  <table width="100%" border="0" cellspacing="0" cellpadding="0">
						<tr>
						  <td class="header-logo" align="left" valign="top" style="padding-left:12px; padding-right:205px;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e369dbf27f65e85d738a4e5c2b5ae9eb560fa6ab671f244e76050d700b959a9da4f8bf2f504512d162b7ee0e74211dfdfe"  target="_blank" style="text-decoration:none"><img style="display:block;color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:16px;" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Order_Confirmation_03.jpg" class="header-logoimg" width="151" height="22" border="0" alt="KOHL'S"/></a></td>
						  <td class="header-mobilenav" align="left" valign="middle">
							<table width="100%" border="0" cellspacing="0" cellpadding="0">
							  <tr>
								<td align="left" valign="middle" style="font-size:0%;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3c6fa82e8ad4f94484e56582dc9c1c8780031ba29ddf949f4991a1b174665c6114312bab2c47d96a31f8cb74206c57be6"  target="_blank" style="text-decoration:none;"><img style="display:block;color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:12px;" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Order_Confirmation_06.jpg" width="96" class="header-myaccount" height="11" border="0" alt="MY ACCOUNT"/></a></td>
								<td align="left" valign="middle" style="padding-right:6px; padding-left:4px; font-size:0%"><img style="display:block;color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:16px;" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Order_Confirmation_08.jpg" width="2" class="header-menuborder" height="14" border="0" alt=""/></td>
								<td align="left" valign="middle"  style="padding-right:12px; font-size:0%;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3c1995d88b09ba6cf2299c8e75df54911b8c427137a4268b35a6005b557da191f02798f29d69736ba08092356c6459201"  target="_blank" style="text-decoration:none"><img style="display:block;color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:12px;" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Order_Confirmation_10.jpg" class="header-customerservice" width="145" height="11" border="0" alt="CUSTOMER SERVICE"/></a></td>
							  </tr>
							</table>
						  </td>
						</tr>
					  </table>
					</td>
				  </tr>
				  <!--Header ends-->

				  <!--Header Menu starts-->
				  <tr>
					<td align="left" valign="top">
					  <table width="100%" border="0" cellspacing="0" cellpadding="0">

					   <tr>
						  <td align="center" valign="top" style="padding-top:15px;">
							<table border="0" cellspacing="0" cellpadding="0" align="center">
							  <tr>
								<td align="center" valign="middle" class="header-menufont header-menupad-first" style="padding:0 20px 0 20px; font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; font-weight:bold; color: #454545;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e301611c1cac3bb6744dfe4ef557f5f45383a1883fdeeb9430a358cbf665e54778af8d3d092e398002fd6ad9e571a158a6"  target="_blank" style="text-decoration:none;color: #454545;"><span style="color: #454545;">FOR THE HOME</span></a></td>
								<td align="center" valign="middle" class="header-menufont header-menupad" style="padding:0 20px 0 20px; font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; font-weight:bold; color: #454545;"> <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3383feec9719dffd855ec3872429918572f97a6132f9e35cfdc384c7a9b67f0bc701ccd16179d882577dccec1da3fffb5"  target="_blank" style="text-decoration:none;color: #454545;"><span style="color: #454545;">WOMEN</span></a> </td>
								<td align="center" valign="middle" class="header-menufont header-menupad" style="padding:0 20px 0 20px; font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; font-weight:bold; color: #454545;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e39b9684edc8a18c7a58ba98a9b97b313f23d8c0a2956db58e22833a3493ed94552b414301161a882383e2d9ebd1ab2bee"  target="_blank" style="text-decoration:none;color: #454545;"><span style="color: #454545;">MEN</span></a></td>
								<td align="center" valign="middle" class="header-menufont header-menupad" style="padding:0 20px 0 20px; font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; font-weight:bold; color: #454545;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3597d8211c4ffc6026a1d7e2faa6d63aa7d81dba0776287877a6d390c5e4eeaaaa087c8c594be8e580f3b5493f3b9b2d0"  target="_blank" style="text-decoration:none;color: #454545;"><span style="color: #454545;">KIDS</span></a></td>
								<td align="center" valign="middle" class="header-menufont header-menupad header-mobile-hidden" style="padding:0 20px 0 20px; font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; font-weight:bold; color: #454545;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e37c9172a3d2b10fc1e187a990d1988dbd608c074ee179eba0ff5409c7b8b6dfc6e33e504e1c5f9105433dcce89eaeda3b"  target="_blank" style="text-decoration:none;color: #454545;"><span style="color: #454545;">BABY</span></a> </td>
								<td align="center" valign="middle" class="header-menufontlast header-menupad-last" style="padding:0 20px 0 20px; font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; font-weight:bold; color: #fd232b;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3a172c0f6ed805d2992af462181b179a86dd932ce535327da2c6dd2d64c589c985653a4d1fbe0654f9fb3c3dc57c7c19c"  target="_blank" style="text-decoration:none;color: #fd232b;"><span style="color: #fd232b;">CLEARANCE</span></a></td>
							  </tr>
							</table>
						  </td>
						</tr>
					  </table>
					</td>
				  </tr>
				  <!--Header Menu ends-->
				  <!--Header ends-->

				  <tr>
					<td align="left" valign="top" class="padbot25">
					  <table width="100%" border="0" cellspacing="0" cellpadding="0">

									<!--Summary starts-->
						<tr>
						  <td align="center" valign="top" bgcolor="#ffffff" style="padding-left:12px; padding-right:12px; padding-top:35px;" class="moborderpad">
							<table width="100%" border="0" cellspacing="0" cellpadding="0">
							  <tr>
								<td align="left" valign="top" class="drop">
								  <table width="100%" border="0" cellspacing="0" cellpadding="0">
									<tr>
									  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; mso-line-height-rule:exactly; font-size: 30px; color: #454545; line-height:30px;" class="fnt20 colorblk alg-center">Order #░░░░░░░░░░</td>
									</tr>
									<tr>
									  <td class="colorblk alg-center" align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545;">Monday, ░░░░░ 30, ░░░░ at 11:░░ AM Central Time</td>
									</tr>
								  </table>
								</td>
								<td align="left" valign="top" width="255" class="mobile-hidden">
								  <table width="100%" border="0" cellspacing="0" cellpadding="0">
									<tr>
									  <td align="center" bgcolor="#95c96f" valign="middle" width="255" height="42" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; color: #ffffff; font-weight:bold;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e333b551788a6a0d0f56b3a20b0ece835b21fd71fd08f1e81d653aa213b60fa2bc83fc9046c1abc38d792ad4f92acbd516"  target="_blank" style="text-decoration:none;color: #ffffff;">MANAGE ORDER</a></td>
									</tr>
								  </table>
								</td>
							  </tr>
							</table>
						  </td>
						</tr>
						<tr>
						  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; mso-line-height-rule:exactly; font-size: 16px; color: #454545; padding-left:12px; padding-right:12px; font-weight:bold; line-height:20px; padding-top:14px;" class="fnt12 alg-center">Thanks for shopping with us, ░░░░░░░░!</td>
						</tr>
						<tr>


						  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; mso-line-height-rule:exactly; font-size: 16px; color: #454545; padding-left:12px; padding-right:12px; line-height:20px;" class="mobbotpadbtn fnt12 alg-center">We'll start putting your order together shortly.<br><br>Once your order is ready for pickup, we'll send you an&nbsp;email.<br />Don't ░░░░ to the store just&nbsp;yet.</td>


						</tr>
						<tr>
						  <td align="center" valign="top">
							<!-- Div used to Hide the Content - Starts-->
							<!--[if !mso]><!-->
							<div class="mobile-show" style="font-size: 0%; max-height: 0; overflow: hidden; display: none;">
							  <table width="180" border="0" cellspacing="0" cellpadding="0">
								<tr>
								  <td align="center" bgcolor="#95c96f" valign="middle" width="255" height="30" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 10px; color: #ffffff; font-weight:bold;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e333b551788a6a0d0f56b3a20b0ece835b21fd71fd08f1e81d653aa213b60fa2bc83fc9046c1abc38d792ad4f92acbd516"  target="_blank" style="text-decoration:none;color: #ffffff;">MANAGE ORDER</a></td>
								</tr>
							  </table>
							</div>
							<!--<![endif]-->
							<!-- Div used to Hide the Content - Ends--></td>
						</tr>
							<!--Summary ends-->





						   <!--Shipping Header starts-->


						<!--Shipping Header starts-->
						<tr>
						  <td align="center" valign="top" bgcolor="#ffffff" style="padding-top:30px;">
							<table width="100%" border="0" cellspacing="0" cellpadding="0">

							  <tr>
								<td align="center" valign="top" bgcolor="#f0f0f0" style="border-top: 2px solid #cccccc; padding-left:12px; padding-right:12px; padding-top:18px; padding-bottom:28px;" class="shipping-mobpad shipping-bordermob">
								  <table width="100%" border="0" cellspacing="0" cellpadding="0">



									<tr>
									  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 16px; padding-bottom:5px; font-weight: bold; color: #000001;" class="shipping-item-fs15 shipping-pickup-mobpickpad">░░░ - Pick Up In Store</td>
									</tr>



									<tr>
									  <td align="left" valign="top" class="shipping-drop">
										<table width="100%" border="0" cellspacing="0" cellpadding="0">
										  <tr>
											<td align="left" valign="top" width="130" style="padding-right:15px;">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">

												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; font-weight: bold; color: #000001;" class="shipping-item-fs15">Pick Up</td>
												</tr>

												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001;">(2 items)</td>
												</tr>
											  </table>
											</td>
											<td align="left" valign="top" width="163" style="padding-right:15px;" class="shipping-zeropadrt">
											  <table class="shipping-widthaddr" width="100%" border="0" cellspacing="0" cellpadding="0">



												<tr>
												  <td align="left" valign="top" class="shipping-fnt12">
													<table width="100%" border="0" cellspacing="0" cellpadding="0">
													  <tr>


														<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; font-weight: bold; color: #000001;" class="shipping-item-fs15" >


														  <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e34558cb8eab4896a43594261fc3d0ee2db112003ef81dbcc5c067fb7b6061e45b5659c89b9b93dd910adb149c6e155361dbbc00e301bca5b9"  target="_blank" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; font-weight: bold; color: #255aec; text-decoration:underline"><span style="color: #255aec; text-decoration:underline">Add Pickup Person</span></a>

														</td>
													  </tr>
													</table>
												  </td>
												</tr>





											  </table>
											</td>
										  </tr>
										</table>
									  </td>

									  <td align="left" valign="top" class="shipping-drop shipping-item-mobpadhdr">
										<table width="100%" border="0" cellspacing="0" cellpadding="0">
										  <tr>



											<td align="left" valign="top" width="294">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; font-weight: bold; color: #000001;" class="shipping-item-fs15">Pickup Location</td>
												</tr>
												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001; padding-bottom:3px;" class="shipping-item-fs15"> ░░░░ TUSCANY ░░<br />░░░░░░, ░░ ░░░░░<br />░░░-░░░-░░░░</td>
												</tr>
											  </table>
											</td>



										  </tr>




										</table>
									  </td>
									</tr>
								  </table>
								</td>
							  </tr>
							</table>
						  </td>
						</tr>
						<!--Shipping Header ends-->




						   <!--Shipping Header ends-->

						<!--Shipping Product List starts-->
						<tr>
						  <td align="center" valign="top" bgcolor="#ffffff" style="padding-top:10px;">
							<table width="100%" border="0" cellspacing="0" cellpadding="0">
							  <tr>
								<td align="center" valign="top" bgcolor="#ffffff" style="padding-left:12px; padding-right:12px;" class="prod-mobpad">
								  <table align="center" width="100%" border="0" cellspacing="0" cellpadding="0">



												   <!--Shipping Product starts-->


									<!--Shipping Product starts-->
									<tr>
									  <td align="left" valign="top" class="product-top" style="padding-top:20px;">
										<table width="100%" border="0" cellspacing="0" cellpadding="0">
										  <tr>
											<td align="left" valign="top" width="145" style="padding-right:15px;" class="product-width130left">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top"><img style="display:block;" src="https://media.kohlsimg.com/is/image/kohls/3223556_Speckled_Dot_White?wid=180&amp;hei=180&amp;op_sharpen=1" class="product-width130" width="132" height="132" border="0" alt="" /></td>
												</tr>
											  </table>
											</td>
											<td align="left" valign="top" width="345" style="padding-right:15px;" class="product-drop product-zeropadrt">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">

		<!-- product name and link -->
												<tr>

												  <td align="left" valign="middle"style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; color: #000001;" class="product-mid-contpad3"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3b6164cf2cbdb14ed856f0ea65cb98ba41a202df17d46e00642ed54d3479bfb157e17075aea2ba94da0c4c0a384dbfcd845ea5139994d79be"  target="_blank" style="text-decoration:underline;color: #000001;"><span style="color: #000001;">Plus Size ░░░░░ & ░░░░░░░ Pleated ░░░░░░ Top</span></a></td>

												</tr>

		<!-- product info (color, size, qty) -->

												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 11px; color: #454545; padding-top:2px;" class="product-fnt12 product-lhmob">Color: Speckled ░░░ ░░░░░, Size: ░░, ░░░: 1</td>
												</tr>


		<!-- product sku -->

												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 10px; color: #454545;" class="product-fnt12 product-lhmob">░░░ # ░░░░░░░░</td>
												</tr>


		<!-- sale/regular price -->

												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545; padding-top:2px; padding-bottom:6px;" class="product-fnt12 product-zeropad product-lhmob"><span style="color:#dc0000; font-weight:bold;"> $24.99</span>SALE $40.░░</td>
												</tr>


		<!-- registry item -->


		<!-- gift receipt included -->


		<!-- gift message -->




		<!-- BOGO -->


		<!-- GWP buy -->


		<!-- PWP buy -->


		<!-- Marketplace item -->


											  </table>
											</td>
											<td align="right" valign="top" width="155" style="padding-right:0px;" class="product-drop product-zeropadrt">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">

		<!-- alert message (item unavailable / only x available) -->


		<!-- your price -->

												<tr>
												  <td align="right" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; color: #000001;" class="product-yppad">Your Price:<span style="font-weight:bold;">$24.99</span></td>

												</tr>


		<!-- gift wrapping -->


											  </table>
											</td>
										  </tr>
										</table>
									  </td>
									</tr>
									<!--Shipping Product ends-->
												   <!--Shipping Product ends-->



												   <!--Shipping Product starts-->


									<!--Shipping Product starts-->
									<tr>
									  <td align="left" valign="top" class="product-top" style="padding-top:20px;">
										<table width="100%" border="0" cellspacing="0" cellpadding="0">
										  <tr>
											<td align="left" valign="top" width="145" style="padding-right:15px;" class="product-width130left">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top"><img style="display:block;" src="https://media.kohlsimg.com/is/image/kohls/3223556_Moss_Geometric_Beige?wid=180&amp;hei=180&amp;op_sharpen=1" class="product-width130" width="132" height="132" border="0" alt="" /></td>
												</tr>
											  </table>
											</td>
											<td align="left" valign="top" width="345" style="padding-right:15px;" class="product-drop product-zeropadrt">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">

		<!-- product name and link -->
												<tr>

												  <td align="left" valign="middle"style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; color: #000001;" class="product-mid-contpad3"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3638f1f43cd8f858fb062560d1c936167547c788e7a5a8091eaf9e150cadfc660a2bc9a3f5a9e551d7574982667e3c455b86567574a02e88b"  target="_blank" style="text-decoration:underline;color: #000001;"><span style="color: #000001;">Plus Size ░░░░░ & ░░░░░░░ Pleated ░░░░░░ Top</span></a></td>

												</tr>

		<!-- product info (color, size, qty) -->

												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 11px; color: #454545; padding-top:2px;" class="product-fnt12 product-lhmob">Color: ░░░░ Geometric Beige, Size: ░░, ░░░: 1</td>
												</tr>


		<!-- product sku -->

												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 10px; color: #454545;" class="product-fnt12 product-lhmob">░░░ # ░░░░░░░░</td>
												</tr>


		<!-- sale/regular price -->

												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545; padding-top:2px; padding-bottom:6px;" class="product-fnt12 product-zeropad product-lhmob"><span style="color:#dc0000; font-weight:bold;"> $24.99</span>SALE $40.░░</td>
												</tr>


		<!-- registry item -->


		<!-- gift receipt included -->


		<!-- gift message -->




		<!-- BOGO -->


		<!-- GWP buy -->


		<!-- PWP buy -->


		<!-- Marketplace item -->


											  </table>
											</td>
											<td align="right" valign="top" width="155" style="padding-right:0px;" class="product-drop product-zeropadrt">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">

		<!-- alert message (item unavailable / only x available) -->


		<!-- your price -->

												<tr>
												  <td align="right" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; color: #000001;" class="product-yppad">Your Price:<span style="font-weight:bold;">$24.99</span></td>

												</tr>


		<!-- gift wrapping -->


											  </table>
											</td>
										  </tr>
										</table>
									  </td>
									</tr>
									<!--Shipping Product ends-->
												   <!--Shipping Product ends-->



								  </table>
								</td>
							  </tr>

							</table>
						  </td>
						</tr>
						<!--Shipping Product List ends-->



						<!--Summary starts-->
						<!--Order Summary starts-->
						<tr>
						  <td align="left" valign="top" bgcolor="#ffffff" style="padding-top:20px;">
							<table width="100%" border="0" cellspacing="0" cellpadding="0">
							  <tr>
								<td align="center" valign="top" style="padding:0 12px 0 12px; border-top: 1px solid #cccccc;">
								  <table width="100%" border="0" cellspacing="0" cellpadding="0">



									<!-- payment info begins -->
									<tr>
									  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 18px; font-weight: bold; color: #454545; padding-top:25px; padding-bottom:18px;" class="summary-item-fs15 summary-colorblk">Payment Information</td>
									</tr>
									<tr>
									  <td align="left" valign="top">
										<table width="100%" border="0" cellspacing="0" cellpadding="0">
										  <tr>



											<!-- payment method begins -->
											<td align="left" valign="top" class="summary-drop summary-item-mobpadbottom">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top" width="250" style="padding-right:65px;" class="summary-padrtzero">
													<table width="100%" border="0" cellspacing="0" cellpadding="0">
													  <tr>
														<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; padding-bottom:5px; font-weight: bold; color: #454545;" class="summary-fnt13">Payment Method</td>
													  </tr>
													  <tr>
														<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001;">
														  <table width="100%" border="0" cellspacing="0" cellpadding="0">



															<tr>
															  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545;">░░░░░░░░░░░ ░░░░░</td>
															  <td align="right" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545;">$48.46</td>
															</tr>



														   </table>
														</td>
													  </tr>
													</table>
												  </td>
												</tr>
											  </table>
											</td>
											<!-- payment method ends -->



											<!-- billing info begins -->
											<td align="left" valign="top" class="summary-drop">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top" width="240" style="padding-right:70px">
													<table width="100%" border="0" cellspacing="0" cellpadding="0">
													  <tr>
														<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 14px; padding-bottom:5px; font-weight: bold; color: #454545;" class="summary-fnt13">Billing Address</td>
													  </tr>
													  <tr>
														<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001;">
														  <table width="100%" border="0" cellspacing="0" cellpadding="0">
															<tr>
															  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545;" class="summary-fnt13">░░░░░░░░ ░░░░░<br>░░░░ Georgetown Dr.<br>░░░░░░, ░░ ░░░░░</td>
															</tr>
														  </table>
														</td>
													  </tr>
													</table>
												  </td>
												</tr>
											  </table>
											</td>
											<!-- billing info ends -->

										  </tr>
										</table>
									  </td>
									</tr>
									<!-- payment info ends -->


									<!-- order summary & earnings begins -->
									<tr>
									  <td align="left" valign="top" style="padding-top:20px;">
										<table width="100%" border="0" cellspacing="0" cellpadding="0">
										  <tr>



											<!-- order summary begins -->
											<td align="left" valign="top" class="summary-drop">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top" width="250">
													<table width="100%" border="0" cellspacing="0" cellpadding="0">
													  <tr>
														<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 18px; padding-bottom:18px; font-weight: bold; color: #454545;" class="summary-item-fs15 summary-colorblk summary-ordersummary-padbtos">Order Summary</td>
													  </tr>
													  <tr>
														<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001;">
														  <table width="100%" border="0" cellspacing="0" cellpadding="0">
															  <tr>
															  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545;">Subtotal:</td>
															  <td align="right" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545;">$49.98</td>
															</tr>





															<tr>
															  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #74b64a;">░░░░░░░░_░░░░░░</td>
															  <td align="right" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #74b64a;"><span style="font-weight:bold;">- $5.░░</span></td>
															</tr>






															 <tr>
															  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545; padding-top:10px;">Shipping:</td>
															  <td align="right" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545; padding-top:10px;"><span style="font-weight:bold; color:#95c96f;">FREE</span></td>
															</tr>


															<tr>
															  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545;">Sale Tax:</td>
															  <td align="right" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #454545;">$3.48</td>
															</tr>
															<tr>
															  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 18px; font-weight:bold; color: #454545; padding-top:10px;">Total:</td>
															  <td align="right" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 18px; color: #454545; font-weight:bold; padding-top:10px;">$48.46</td>
															</tr>
														  </table>
														</td>
													  </tr>
													</table>
												  </td>
												</tr>
											  </table>
											</td>
											<!-- order summary ends -->


											<!-- earnings & savings begins -->
											<td align="left" valign="top" style="padding-left:65px;" class="summary-drop summary-item-mobpadhdr">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top" width="310" style="padding-right:0px">
													<table width="100%" border="0" cellspacing="0" cellpadding="0">


													  <!-- your earnings begins -->
													  <tr>
														<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 18px; padding-bottom:18px; font-weight: bold; color: #454545;" class="summary-item-fs15 summary-colorblk summary-earning-padbtye">Your Earnings</td>
													  </tr>



													  <tr>
														<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001;">
														  <table width="100%" border="0" cellspacing="0" cellpadding="0">
															<tr>
															  <td align="left" valign="top" width="45"><img style="display:block;color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:16px;" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Order_Confirmation_40.jpg" width="45" height="23" border="0" alt="" /></td>
															  <td align="left" valign="top">
																<table width="100%" border="0" cellspacing="0" cellpadding="0">
																  <tr>
																	<td align="left" valign="top">
																	  <table width="100%" border="0" cellspacing="0" cellpadding="0">
																		<tr>
																		  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:14px; mso-line-height-rule:exactly; font-size: 12px; color: #666666; font-weight:bold; padding-left:5px;">░░░░░░░ Rewards</td>

																		  <td align="right" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:14px; mso-line-height-rule:exactly; font-size: 12px; color: #666666; font-weight:bold;">45 potential points</td>


																		</tr>
																	  </table>
																	</td>
																  </tr>


																  <tr>
																	<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:12px; mso-line-height-rule:exactly; font-size: 10px; color: #666666; padding-left:5px;">Your points ░░░░ appear in your account up to 48 hours after this order is
		marked Complete in your Purchase History.</td>
																  </tr>


																</table>
															  </td>
															</tr>
														  </table>
														</td>
													  </tr>
													  <!-- your earnings ends -->



													  <!-- your savings begins -->
													  <tr>
														<td width="270" align="center" valign="bottom" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; padding-top:15px;" class="summary-yrstp">
														  <table width="270" border="0" cellspacing="0" cellpadding="0">
															<tr>
															  <td width="26" align="left" valign="top" style="font-size:0%"><img style="display:block;color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:16px;" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Order_Confirmation_81.jpg" width="26" height="52" border="0" alt="" /></td>
															  <td width="218" align="left" valign="top">
																<table width="218" border="0" cellspacing="0" cellpadding="0">
																  <tr>
																	<td width="218" align="left" valign="top" style="font-size:0%"><img style="display:block;color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:16px;" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Order_Confirmation_82.jpg" width="218" height="17" border="0" alt="" /></td>
																  </tr>
																  <tr>
																	<td width="218" align="center" height="22" valign="middle" style="font-family: Arial, Helvetica, sans-serif;font-size: 18px; color: #666666; font-weight:bold;" class="summary-yrs">Your Savings $35.░░</td>
																  </tr>
																  <tr>
																	<td width="218" align="left" valign="top"><img style="display:block;color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:16px;" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Order_Confirmation_84.jpg" width="218" height="13" border="0" alt="" /></td>
																  </tr>
																</table>
															  </td>
															  <td width="26" align="left" valign="top" style="font-size:0%"><img style="display:block;color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:16px;" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Order_Confirmation_83.jpg" width="26" height="52" border="0" alt="" /></td>
															</tr>
														  </table>
														</td>
													  </tr>
													  <!-- your savings ends -->

													</table>
												  </td>
												</tr>
											  </table>
											</td>
											<!-- earnings & savings ends -->


										  </tr>
										</table>
									  </td>
									</tr>
									<!-- order summary & earnings ends -->

								  </table>
								</td>
							  </tr>
							</table>
						  </td>
						</tr>
						<!--Order Summary ends-->
						<!--Summary ends-->

						<!--Return policy starts-->
						<tr>
							  <td align="left" valign="top" bgcolor="#ffffff" style="padding-top:20px;">
								<table width="100%" border="0" cellspacing="0" cellpadding="0">
									<tr>
										  <td align="center" valign="top" bgcolor="#ffffff" style="padding-left:12px; padding-right:12px; border-top: 1px solid #cccccc;">
											<table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; mso-line-height-rule:exactly; font-size: 16px; color: #454545; font-weight:normal; line-height:20px; padding-top:20px;" class="fnt12 alg-center">Please visit<a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e347bccbdb1057167a508c6511b0ba2bc78a0b6620c64a7d31cffc00544be35c5f0eaaa3bee47f841481d99bbd84d67785"  target="_blank" style="text-decoration:underline;color:#000001;">░░░░░.░░░/returns</a> to learn about our hassle-free return policy. Some items, such as premium
		electronics, must be returned within 30 days.</td>
												</tr>
											</table>
										</td>
									</tr>
								</table>
							</td>
						</tr>
						<!--Return policy ends-->


						<!--Recommendations starts-->


						<!--Recommendations starts-->
						<tr>
						  <td align="left" valign="top" bgcolor="#ffffff" style="padding-top:20px;">
							<table width="100%" border="0" cellspacing="0" cellpadding="0">
									<tr>

									  <td align="center" valign="top" bgcolor="#ffffff" class="recs-nopadtp" style="padding-left:12px; padding-right:12px; border-top: 1px solid #cccccc;">

										<table width="100%" border="0" cellspacing="0" cellpadding="0">
										  <tr>
											<td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:20px; mso-line-height-rule:exactly; font-size: 18px; padding-top: 25px; font-weight: bold; color: #000001;" class="recs-fnt19 recs-alg-center">Here are some great items from your store:</td>
											</tr>
												<tr>
												  <td align="center" valign="top"><img class="recs-mobile-hidden" usemap="#rrImgMap_desktop" alt="Product Recommendations" src="http://rm.recs.richrelevance.com/rrmail/imgreq?a=648c894ab44bc04a&amp;cpi=order_conf&amp;userId=5571792788&amp;seedProductIds=3223556|3223556&amp;zoneName=responsive_email&amp;date=04/30/2018&amp;version=3&amp;ln=email_desktop" width="616" border="0" style="display: block;"></td>
												</tr>
												<tr>
																	<td valign="top" align="center" bgcolor="#ffffff" style="font-size:0px; mso-line-height-rule:exactly; line-height:0px;">
																		<!--[if !mso]><!--><div class="recs-mobile-show" style="display:none; width:0px; max-height:0px; overflow:hidden;">
																			<table border="0" cellpadding="0" cellspacing="0" width="100%">
																				<tr><td valign="top" align="center" bgcolor="#ffffff"><img class="recs-header-img" usemap="#rrImgMap_mobile" alt="Product Recommendations" src="http://rm.recs.richrelevance.com/rrmail/imgreq?a=648c894ab44bc04a&amp;cpi=order_conf&amp;userId=5571792788&amp;seedProductIds=3223556|3223556&amp;zoneName=responsive_email&amp;date=04/30/2018&amp;version=3&amp;ln=email_mobile" width="290" border="0" style="width:290px; height:auto;"></td></tr>
																			</table>
																		</div><!--<![endif]-->
																	</td>
												</tr>

												<map name="rrImgMap_desktop" id="rrImgMap_desktop">
																  <area shape="rect" coords="0,11,146,217" alias="Product Recommendations 1" target="_blank" href="http://rm.recs.richrelevance.com/rrmail/rmclick?a=648c894ab44bc04a&cpi=order_conf&userId=5571792788&seedProductIds=3223556|3223556&zoneName=responsive_email&date=04/30/2018&version=3&pfm=recs-order_conf_rr&ln=email_desktop&lid=1" />
																  <area shape="rect" coords="154,12,303,218" alias="Product Recommendations 2" target="_blank" href="http://rm.recs.richrelevance.com/rrmail/rmclick?a=648c894ab44bc04a&cpi=order_conf&userId=5571792788&seedProductIds=3223556|3223556&zoneName=responsive_email&date=04/30/2018&version=3&pfm=recs-order_conf_rr&ln=email_desktop&lid=2" />
																  <area shape="rect" coords="313,11,459,219" alias="Product Recommendations 3" target="_blank" href="http://rm.recs.richrelevance.com/rrmail/rmclick?a=648c894ab44bc04a&cpi=order_conf&userId=5571792788&seedProductIds=3223556|3223556&zoneName=responsive_email&date=04/30/2018&version=3&pfm=recs-order_conf_rr&ln=email_desktop&lid=3" />
																  <area shape="rect" coords="469,11,616,217" alias="Product Recommendations 4" target="_blank" href="http://rm.recs.richrelevance.com/rrmail/rmclick?a=648c894ab44bc04a&cpi=order_conf&userId=5571792788&seedProductIds=3223556|3223556&zoneName=responsive_email&date=04/30/2018&version=3&pfm=recs-order_conf_rr&ln=email_desktop&lid=4" />
																</map>
																<map name="rrImgMap_mobile" id="rrImgMap_mobile">
																  <area shape="rect" coords="3,12,129,221" alias="Product Recommendations Mobile 1" target="_blank" href="http://rm.recs.richrelevance.com/rrmail/rmclick?a=648c894ab44bc04a&cpi=order_conf&userId=5571792788&seedProductIds=3223556|3223556&zoneName=responsive_email&date=04/30/2018&version=3&pfm=recs-order_conf_rr&ln=email_mobile&lid=1" />
																  <area shape="rect" coords="161,12,301,220" alias="Product Recommendations Mobile 2" target="_blank" href="http://rm.recs.richrelevance.com/rrmail/rmclick?a=648c894ab44bc04a&cpi=order_conf&userId=5571792788&seedProductIds=3223556|3223556&zoneName=responsive_email&date=04/30/2018&version=3&pfm=recs-order_conf_rr&ln=email_mobile&lid=2" />
																  <area shape="rect" coords="0,220,129,436" alias="Product Recommendations Mobile 3" target="_blank" href="http://rm.recs.richrelevance.com/rrmail/rmclick?a=648c894ab44bc04a&cpi=order_conf&userId=5571792788&seedProductIds=3223556|3223556&zoneName=responsive_email&date=04/30/2018&version=3&pfm=recs-order_conf_rr&ln=email_mobile&lid=3" />
																  <area shape="rect" coords="163,219,287,430" alias="Product Recommendations Mobile 4" target="_blank" href="http://rm.recs.richrelevance.com/rrmail/rmclick?a=648c894ab44bc04a&cpi=order_conf&userId=5571792788&seedProductIds=3223556|3223556&zoneName=responsive_email&date=04/30/2018&version=3&pfm=recs-order_conf_rr&ln=email_mobile&lid=4" />
																</map>
															</table>
														</td>
													</tr>
							</table>
						  </td>
						</tr>
						<!--Recommendations ends-->

						<!--Recommendations ends-->



						<!--Help starts-->


						<!--Help starts-->
						<tr>
						  <td align="left" valign="top" bgcolor="#ffffff" style="padding-top:20px;">
							<table width="100%" border="0" cellspacing="0" cellpadding="0">
							  <tr>
								<td align="center" valign="top" bgcolor="#ffffff" class="help-nopadtp help-nhpad" style="padding-left:12px; padding-right:12px; border-top: 1px solid #cccccc;">
								  <table class="help-width288" width="100%" border="0" cellspacing="0" cellpadding="0">
									<tr>
									  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:24px; mso-line-height-rule:exactly; font-size: 24px; padding-bottom: 18px;padding-top: 25px; font-weight: bold; color: #000001;" class="help-fnt19">Need help? No problem.</td>
									</tr>
									<tr>
									  <td align="center" valign="top">
										<table width="100%" border="0" cellspacing="0" cellpadding="0">
										  <tr>
											<td align="left" valign="top" class="help-drop">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:18px; mso-line-height-rule:exactly; font-size: 18px; font-weight: bold; color: #000001;" class="help-fnt16">Frequently Asked Questions</td>
												</tr>

												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001; padding-top:3px;" class="help-fnt13"><a href="http://click.t.kohls.com/?qs=56503053baa0b4710a547700463eefdf4ef6f5f75419d1d20cd6851cfc971b3d5cf977f1b4c0198c2b36d3cf2fdde42b7de731108e7f896240ee51a4d9cfe43d"  target="_blank" style="text-decoration:underline;color:#000001;"><span style="color:#000001;">How do I track my order?</span></a></td>
												</tr>


												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001; padding-top:3px;" class="help-fnt13"><a href="http://click.t.kohls.com/?qs=56503053baa0b4711ae6c8fb4368472f35caf851b4091899a5d2c1026ed3dcf81941f1652b015bc681288396e1b91926a6921d4f0a48fe4e3464cdf1ee9d0b83"  target="_blank" style="text-decoration:underline;color:#000001;"><span style="color:#000001;">Can I ░░░░░░ or change my order?</span></a></td>
												</tr>


												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001;padding-top:3px;" class="help-fnt13"><a href="http://click.t.kohls.com/?qs=56503053baa0b4717c47f6889586dd64f9b1c587c6f282ef00157455ca35e827f09e0dc0a8e2a45fe1ab2f1a0af2ba62872293bea227c1e39de1f8d552c42b6b"  target="_blank" style="text-decoration:underline;color:#000001;"><span style="color:#000001;">Can I add a pickup person to my order?</span></a></td>
												</tr>


												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001;padding-top:3px;" class="help-fnt13"><a href="http://click.t.kohls.com/?qs=56503053baa0b471d170b168a58527949c31d88ff36d0dd73379e7c3b84997e28f26fb4dcc678d8c24e7140402cb05ea790c19f66c30cf59c9a3c1d4c1a615f9"  target="_blank" style="text-decoration:underline;color:#000001;"><span style="color:#000001;">When ░░░░ my order arrive?</span></a></td>
												</tr>


												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001;padding-top:3px;" class="help-fnt13"><a href="http://click.t.kohls.com/?qs=56503053baa0b4717690979a6463579961a6db2f051e0b5ba55a0ef97a02c57c456d7e1815357c29840302f31256ecaebb9639f3aa5a5d9b70c1c705b3dc015a"  target="_blank" style="text-decoration:underline;color:#000001;"><span style="color:#000001;">When ░░░░ I be charged?</span></a></td>
												</tr>

											  </table>
											</td>
											<td align="left" valign="top" class="help-drop help-padgpick" width="185">
											  <table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:18px; mso-line-height-rule:exactly; font-size: 18px; font-weight: bold; color: #000001;" class="help-fnt16">Get more information</td>
												</tr>
												<tr>
												  <td align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001;padding-top:3px;" class="help-fnt13"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3dd58e7568ecd0043915c5b2b13542d4fbb90ea82fd98617c0300075ca0eba09557ab23f18a9ebd6b0a1d4b3909a5f7e9d7d984fa3cbb3938"  target="_blank" style="text-decoration:underline;color:#000001;"><span style="color:#000001;">Customer Service</span></a></td>
												</tr>
												<tr>
												  <td align="left" valign="top" style="padding-top:12px;">
													<table width="130" border="0" cellspacing="0" cellpadding="0">
													  <tr>
														<td height="32" width="130" align="center" valign="middle" style="font-family: Arial, Helvetica, sans-serif; line-height:16px; mso-line-height-rule:exactly; font-size: 12px; color: #000001; border: 1px solid #000001;" class="help-fnt13"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e328504b4d31440e3764efb0fab517d90202f72785185a271040f8c2844c2ea16351041a41bee5f3519a3fcc130644935c10d76eef6c07b603"  target="_blank" style="text-decoration:none;color:#000001;font-weight:bold;"><span style="color:#000001;">CONTACT US</span></a></td>
													  </tr>
													</table>
												  </td>
												</tr>
											  </table>
											</td>
										  </tr>
										</table>
									  </td>
									</tr>
								  </table>
								</td>
							  </tr>
							</table>
						  </td>
						</tr>
						<!--Help ends-->
						<!--Help ends-->




						<!-- LiveIntent -->

						<!--start-->
						<tr>

							  <td align="left" valign="top" bgcolor="#ffffff" style="padding-top:20px;">
								<table width="100%" border="0" cellspacing="0" cellpadding="0">
									<tr>
										  <td align="center" valign="top" bgcolor="#ffffff" class="ads-nopadtp" style="padding-left:10px; padding-right:10px; border-top: 1px solid #cccccc;">

											<table width="100%" border="0" cellspacing="0" cellpadding="0">


												  <tr>
													<td align="center" valign="top">
														<table width="100%" border="0" cellspacing="0" cellpadding="0">
															  <tr>

																  <td align="center" valign="top" class="ads-drop" style="padding-top: 18px;">
																	<table width="100%" border="0" cellspacing="0" cellpadding="0">
																		<tr>
																			<td align="center" valign="top">

																				<table width="300" border="0" cellpadding="0" cellspacing="0" >
																					<tr>
																						<td colspan="2"><a style="display: block; width: 300px; height: 250px;" href="http://click.t.kohls.com/?qs=56503053baa0b47150aa79a3e410b063c0259168535db2c239e58c6611e9bafca1ecb5c09589d9e029c04308c7754501e9173e883b9c68323ce56293a391297f" rel="nofollow" ><img src="http://li.kohls.com/imp?s=207249&sz=300x250&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" border="0" width="300" height="250"/></a></td>
																					</tr>
																					<tr style="display:block; height:1px; line-height:1px;">
																						<td><img src="http://li.kohls.com/imp?s=207250&sz=1x1&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" height="1" width="10" /></td>
																						<td><img src="http://li.kohls.com/imp?s=207251&sz=1x1&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" height="1" width="10" /></td>
																					</tr>
																					<tr>
																						<td align="left"><a href="http://click.t.kohls.com/?qs=56503053baa0b4719cc7208d1784b56fb3b90240ef89fd739c655d9f6ea1a3018336059ee1daba815a3b928fac832f9ffb98f47130abf9245e899ec0fae2dde7" rel="nofollow"><img src="http://li.kohls.com/imp?s=119386&sz=116x15&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" border="0"/></a></td>
																						<td align="right"><a href="http://click.t.kohls.com/?qs=56503053baa0b471cf9f158a4977b6d61b8b8b8d1dfc6f28b201b5eb7d8c20c5f30f1893a332f890ebcc5d9dea93bd171f951531d6e80e50f88c6bd2151c47ed" rel="nofollow"><img src="http://li.kohls.com/imp?s=119387&sz=69x15&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" border="0"/></a></td>
																					</tr>
																				</table>

																			</td>
																		</tr>
																	</table>
																</td>

																  <td align="center" valign="top" class="ads-drop" style="padding-top: 18px; padding-left:20px;">
																	<table width="100%" border="0" cellspacing="0" cellpadding="0">
																		<tr>
																			<td align="center" valign="top">

																				<table width="300" border="0" cellpadding="0" cellspacing="0" >
																					<tr>
																						<td colspan="2"><a style="display: block; width: 300px; height: 250px;" href="http://click.t.kohls.com/?qs=56503053baa0b4715a849ecd755a0c1ba63faa96261dc7b31a596ea89445c6f4fccadcf504f0144f25b4db84625a3427736fd193f8cbaa3742a6f2917352a27c" rel="nofollow" ><img src="http://li.kohls.com/imp?s=207252&sz=300x250&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" border="0" width="300" height="250"/></a></td>
																					</tr>
																					<tr style="display:block; height:1px; line-height:1px;">
																						<td><img src="http://li.kohls.com/imp?s=207253&sz=1x1&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" height="1" width="10" /></td>
																						<td><img src="http://li.kohls.com/imp?s=207254&sz=1x1&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" height="1" width="10" /></td>
																					</tr>
																					<tr>
																						<td align="left"><a href="http://click.t.kohls.com/?qs=56503053baa0b471d5534558b1820d0821bf917dcc242dc8e21900a60de02a28c0bd1f18a8380ea5933f2ad4ec9962a3f3d28fa993b28e64b09d9c0eddb7431b" rel="nofollow"><img src="http://li.kohls.com/imp?s=119386&sz=116x15&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" border="0"/></a></td>
																						<td align="right"><a href="http://click.t.kohls.com/?qs=56503053baa0b471e8b1e461bc2be0e40bba3271b87c98273ef88e3fdb4ceddb7224213b04441c0cf208b09936635849643ab9750a8e8f2ff00a2d3ee0d668c0" rel="nofollow"><img src="http://li.kohls.com/imp?s=119387&sz=69x15&li=&m=c217bdff9dbdac53ab15c5941af7703a&p=04302018100824356" border="0"/></a></td>
																					</tr>
																				</table>

																			</td>
																		</tr>
																	</table>
																</td>


															  </tr>
														</table>
													</td>
												  </tr>


											</table>
										  </td>
									</tr>
								</table>
							  </td>
						</tr>

						<!--end-->


						<!--Footer Starts-->
						<!--Footer Starts-->
		<tr>
		  <td align="center" valign="top" style="padding-top:25px; padding-bottom:25px;">
			<table width="100%" border="0" cellspacing="0" cellpadding="0">
			  <tr>
				<td align="center" valign="top" style="padding-top:25px; padding-bottom:25px; background-color:#f0f0f0;">
				  <table width="100%" border="0" cellspacing="0" cellpadding="0">
					<tr>
					  <td align="center" valign="top">
						<table width="100%" border="0" cellspacing="0" cellpadding="0">
						  <tr>
							<td align="center" valign="top" class="footer-drop">
							  <table class="footer-width90" align="center" width="100%" border="0" cellspacing="0" cellpadding="0">
								<tr>
								  <td class="footer-width-auto" align="center" valign="top" style="font-family: Arial, Helvetica, sans-serif; line-height:18px; mso-line-height-rule:exactly; font-size:14px; color:#454545; font-weight:normal;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e336ae4bfaa59f38bc0bc3a67d1731d52d87f4f687490b59b161ed10e22d7b0196410d4d296df17bf42cf63b570882076b"  target="_blank" style="text-decoration:none; color:#454545;">░░░░░░&nbsp;CHARGE</a></td>
								  <td class="footer-width-auto footer-noborder" align="center" valign="top" style="border-left: 1px solid #d6d9db; font-family: Arial, Helvetica, sans-serif; line-height:18px; mso-line-height-rule:exactly; font-size: 14px; color:#454545; font-weight:normal;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3c1995d88b09ba6cf2299c8e75df54911b8c427137a4268b35a6005b557da191f02798f29d69736ba08092356c6459201"  target="_blank" style="text-decoration:none; color:#454545;">CUSTOMER&nbsp;SERVICE</a></td>
								</tr>
							  </table>
							</td>
							<td align="center" valign="top" class="footer-drop footer-toppad">
							  <table class="footer-width90" align="center" width="100%" border="0" cellspacing="0" cellpadding="0">
								<tr>
								  <td class="footer-width-auto footer-noborder" align="center" valign="top" style="border-left: 1px solid #d6d9db; font-family: Arial, Helvetica, sans-serif; line-height:18px; mso-line-height-rule:exactly; font-size:14px; color:#454545; font-weight:normal;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e32210ef62fe2c162f5ef963be3db40a520c0b4e3a6aaa2c265e26489a2f79ef5a36cf9116924be36bc914b788644da8be"  target="_blank" style="text-decoration:none; color:#454545;">FIND&nbsp;A&nbsp;STORE</a></td>
								  <td class="footer-width-auto footer-noborder" align="center" valign="top" style="border-left: 1px solid #d6d9db; font-family: Arial, Helvetica, sans-serif; line-height:18px; mso-line-height-rule:exactly; font-size:14px; color:#454545; font-weight:normal;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3498dc34359d1399baf4942cd9a0ac3ac5387445a7b6fe215490b4184a607554da935d33a956a6a1441df28b09ba92ddf"  target="_blank" style="text-decoration:none; color:#454545;">░░░░░░&nbsp;COUPONS</a></td>
								</tr>
							  </table>
							</td>
						  </tr>
						</table>
					  </td>
					</tr>
					<tr>
					  <td align="center" valign="top" style="padding-top:25px;">
						<table class="footer-fullwidth" width="450" border="0" cellspacing="0" cellpadding="0">
						  <tr>
							<td align="center" valign="top">
							  <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3ee92310f64350472936136a1e4449bf4574659077b9b4e249032b5bdd1fe4c9107f230ccf3232f9dd03d11ce2718e3ee" target="_blank"  style="text-decoration:none;"><img alt="Facebook" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Shared_Footer_Social_FB.jpg" width="26" height="26" border="0" style="display:block;"></a>
							</td>
							<td align="center" valign="top">
							  <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3dfa2e00d2f904027ea644fe2d7cac99cb52978352d376a5b78188c45dc81d9c2588d4d9626684d095dfb5380fc9dc83b" target="_blank"  style="text-decoration:none;"><img alt="Twitter" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Shared_Footer_Social_TW.jpg" width="26" height="26" border="0" style="display:block;"></a>
							</td>
							<td align="center" valign="top">
							  <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3807b1df705526579257502935567fffa0aa4a5ca69c2aff83727e079e6d5a061f9c6b3d6e5c46a0c118be76a2e45a9f4" target="_blank"  style="text-decoration:none;"><img alt="Google" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Shared_Footer_Social_GP.jpg" width="33" height="26" border="0" style="display:block;"></a>
							</td>
							<td align="center" valign="top">
							  <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3b303dc19761889273544fdd1a4aff63c1f5af2b3d2bbfe4801a240b1c200ebb278e96b069cba02ebf1d08137aa2b258a" target="_blank"  style="text-decoration:none;"><img alt="Pinterest" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Shared_Footer_Social_PT.jpg" width="26" height="26" border="0" style="display:block;"></a>
							</td>
							<td align="center" valign="top">
							  <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3e7fbf31ff24b1336f28ec2f76d234cc5418daeed01e4fef2c4b743399bfcccddfead2737aba12ac5fcb110731f65cb53" target="_blank"  style="text-decoration:none;"><img alt="Instagram" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Shared_Footer_Social_IN.jpg" width="26" height="26" border="0" style="display:block;"></a>
							</td>
							<td align="center" valign="top">
							  <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e34bfbbd192fc219f426f13ef63b3d9f7814805cd11bbf8fd5e129c57f2d58e4169219abf2d5631b5b7c0b16d57356495f" target="_blank"  style="text-decoration:none;"><img alt="YouTube" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Shared_Footer_Social_YT.jpg" width="26" height="26" border="0" style="display:block;"></a>
							</td>
						  </tr>
						</table>
					  </td>
					</tr>
					<tr>
					  <td align="center" valign="top" style="padding-top:25px;">
						<table width="300" border="0" cellspacing="0" cellpadding="0">
						  <tr>
							<td align="center" valign="top">
							  <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e35332846fa46e5e7a598ab37eaec6bb39191df3180cb9435affb139c73e29f53fad995867754e65c8a37452c214a1d0a0" target="_blank"  style="text-decoration:none;"><img alt="Download on the App Store" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Shared_Footer_App_Store.jpg" width="120" height="35" border="0" style="display:block;"></a>
							</td>
							<td align="center" valign="top">
							  <a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e383ffa54bf921f63ca7a6eb01bc40ece6b87272d0e57358daa7faddc763bc69f9d115db5909192d5c3e4c67635eabac80" target="_blank"  style="text-decoration:none;"><img alt="Get it on Google Play" src="http://image.t.kohls.com/lib/fe8a1272716d007476/m/1/Shared_Footer_Google_Play.jpg" width="114" height="35" border="0" style="display:block;"></a>
							</td>
						  </tr>
						</table>
					  </td>
					</tr>
					<tr>
					  <td align="center" valign="top" style="padding: 15px 20px 0 20px; font-family: Arial, Helvetica, sans-serif; line-height:15px; mso-line-height-rule:exactly; font-size: 11px; color:#454545; font-weight:normal;"> Android, Google Play and the Google Play logo are trademarks of Google
		Inc.<br>App Store is a service ░░░░ of Apple Inc.</td>
					</tr>
					<tr>
					  <td align="center" valign="top" style="padding-top:25px;">
						<table class="footer-width90" align="center" width="550" border="0" cellspacing="0" cellpadding="0">
						  <tr>
							<td align="center" valign="top" style="padding:0 10px 0 10px; font-family: Arial, Helvetica, sans-serif; line-height:18px; mso-line-height-rule:exactly; font-size:14px; color:#454545; font-weight:normal;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e354c07799605b750e874ec36c8262ec2186c2cd9454d441e8e95c2acc437fbb68e8bffec128efcda4e8078420055cb6d7"  target="_blank" style="text-decoration:none; color:#454545;">SECURITY&amp; PRIVACY POLICY</a></td>
							<td align="center" valign="top" style="padding:0 10px 0 10px; border-left: 1px solid #d6d9db; font-family: Arial, Helvetica, sans-serif; line-height:18px; mso-line-height-rule:exactly; font-size: 14px; color:#454545; font-weight:normal;"><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e31fa77734f4adb10d84d93767d1989dbbeeda3f903afc8892b60f57f768697a8c33e33a8c3332aa2d9ab9681cd5482e92"  target="_blank" style="text-decoration:none; color:#454545;">HASSLE-FREE RETURNS POLICY</a></td>
						  </tr>
						</table>
					  </td>
					</tr>
				  </table>
				</td>
			  </tr>
			</table>
		  </td>
		</tr>
		<!--Footer ends-->

		<!--Tracking starts-->
		<tr>
		  <td align="center" valign="top" bgcolor="#ffffff" class="footer-nopadtp" style="padding-left:12px; padding-right:12px; font-size:0%">
			<table cellpadding="0" cellspacing="0" border="0" width="40" height="6" style="display:none;">
			<tbody>
				<tr>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713600&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713601&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713602&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713603&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713604&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713605&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713606&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713607&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713608&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713609&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713610&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713611&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713612&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713613&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713614&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713615&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713616&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713617&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713618&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
					<td style="font-size:0%"><img src=http://li.kohls.com/imp?s=123713619&sz=2x1&li=Order_Confirmation_Responsive_Dynamic&e=jendowns@yahoo.com&p=411413 width="2" height="6" border="0" /></td>
				</tr>
			</tbody>
		</table>
		  </td>
		</tr>
		<tr>
		  <td align="center" valign="top" bgcolor="#ffffff" class="footer-nopadtp" style="padding-left:12px; padding-right:12px; font-size:0%">

		<img width='0' height='0' src='https://log.dmtry.com/redir/1/1289/3811/_order_confirmation_responsive_dynamic/0/1/157/0/966/1.ver?at=i&d=PxImp&img=1'>
		  </td>
		</tr>
		 <tr>
		  <td align="center" valign="top" bgcolor="#ffffff" class="footer-nopadtp" style="padding-left:12px; padding-right:12px; font-size:0%">

		<img src="https://pippio.com/api/sync?pid=5322&_=1&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;"><img src="https://pippio.com/api/sync?pid=5322&_=2&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;"><img src="https://pippio.com/api/sync?pid=5322&_=3&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;"><img src="https://pippio.com/api/sync?pid=5322&_=4&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;"><img src="https://pippio.com/api/sync?pid=5322&_=5&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;"><img src="https://pippio.com/api/sync?pid=5322&_=6&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;"><img src="https://pippio.com/api/sync?pid=5322&_=7&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;"><img src="https://pippio.com/api/sync?pid=5322&_=8&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;"><img src="https://pippio.com/api/sync?pid=5322&_=9&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;"><img src="https://pippio.com/api/sync?pid=5322&_=10&it=4&iv=c217bdff9dbdac53ab15c5941af7703a&it=4&iv=c20684712fe7605864e2ff1d571acf274dce3f8a&it=4&iv=c6ce1fe190c9659cd979d6dae2e8e5328da459dfac501b6f9487a3d5f34117fd" width="1"
		height="1" border="0" style="display:none; overflow:hidden;">
		  </td>
		</tr>
		<!--Tracking ends-->
						  <!--Footer ends-->

						  <!--Subfooter Starts-->
								   <tr>
									  <td class="alg-center appleLinks" align="left" valign="top" style="font-family: Arial, Helvetica, sans-serif; font-size: 10px; color: #000001; padding: 0 12px 20px 12px;" class="footer-pad">For other inquiries, email us at<a href="mailto:myaccount.help@kohls.com" alias="Email Us" target="_blank" style="text-decoration:underline;color:#000001;">░░░░░░░░░.help@░░░░░.░░░</a>, or write us at ░░░░░░ Department Stores, Attention: Customer Service, ░░░
		░░░░░░ ░░░░░░░ Drive, ░░░░░░░░░ ░░░░░, ░░ ░░░░░.<br><a href="http://click.t.kohls.com/?qs=3ac93bbccd2308e3d3e647cbd8aa803b49be00b217ff1116cfbca0a3cfe68aedef9caa46e67b6241a1d65707dfb61a6fa89f654a43fffe7d11235722653af4f2"  target="_blank" style="text-decoration:underline;color:#000001;">Find a ░░░░░░</a> ░░░░ you!</td>
										</tr>
										<!--Subfooter Ends-->

					  </table>
					</td>
				  </tr>

				</table>
				<!-- END Table use to set width of email -->
			  </div>
			</td>
		  </tr>
		</table>
		<!-- END Table used to center email Ends-->
		</body>
		</html>

		`,
	})
	if nil != setContentResult.Err {
		log.Fatalf("%+v", setContentResult.Err)
	}

	// Create a channel to receive the screenshot data generated by the
	// event handler.
	results := make(chan *page.CaptureScreenshotResult)

	// Create an event handler that executes when the page load event is
	// received.
	loadEventHandler := socket.NewEventHandler(
		"Page.loadEventFired",

		// This function will generate a screenshot and write the data
		// to the results channel.
		func(response *socket.Response) {

			// Set the device emulation parameters.
			overrideResult := <-tab.Emulation().SetDeviceMetricsOverride(
				&emulation.SetDeviceMetricsOverrideParams{
					Width:  1440,
					Height: 1440,
					ScreenOrientation: &emulation.ScreenOrientation{
						Type:  emulation.OrientationType.PortraitPrimary,
						Angle: 90,
					},
				},
			)
			if nil != overrideResult.Err {
				panic(overrideResult.Err)
			}

			// Capture a screenshot of the current state of the current
			// page.
			screenshotResult := <-tab.Page().CaptureScreenshot(
				&page.CaptureScreenshotParams{
					Format:  page.Format.Jpeg,
					Quality: 50,
				},
			)
			if nil != screenshotResult.Err {
				panic(screenshotResult.Err)
			}

			results <- screenshotResult
		},
	)

	// Add the handler to the stack for this tab. Multiple handlers can
	// listen to the same event.
	tab.AddEventHandler(loadEventHandler)

	// Wait for the handler to fire
	result := <-results

	// Decode the base64 encoded image data
	data, err := base64.StdEncoding.DecodeString(result.Data)
	if nil != err {
		panic(err)
	}

	// Write the generated image to a file
	err = ioutil.WriteFile("/tmp/test/example.jpg", data, 0644)
	if nil != err {
		panic(err)
	}

	fmt.Println("Finished rendering example.jpg")
}
