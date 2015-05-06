// 获取公众平台的文档(http://mp.weixin.qq.com/wiki/home/index.html)
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// http://mp.weixin.qq.com/wiki/home/index.html 左边菜单的数据结构
type Menu struct {
	XMLName    struct{} `xml:"xml" json:"-"`
	Level1List []Level1 `xml:"div"`
}

// 一级目录
type Level1 struct {
	H5 struct {
		Span string `xml:"span"`
		Text string `xml:",chardata"`
	} `xml:"h5"`
	Div struct {
		Level2List []Level2 `xml:"ul>li"`
	} `xml:"div"`
}

// 二级目录
type Level2 struct {
	Anchor struct {
		Href string `xml:"href,attr"`
		Text string `xml:",chardata"`
	} `xml:"a"`
}

// 下载二级目录的html文件
func htmlDownload(dir, filename, url string) (err error) {
	fp := filepath.Join(dir, filename+".html")
	file, err := os.Create(fp)
	if err != nil {
		return
	}
	defer func() {
		file.Close()
		if err != nil {
			os.Remove(fp)
		}
	}()

	httpResp, err := http.Get(url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	_, err = io.Copy(file, httpResp.Body)
	return
}

func main() {
	// 解析xml
	var mn Menu
	if err := xml.Unmarshal(src, &mn); err != nil {
		fmt.Println(err)
		return
	}

	// 创建目标目录
	exePath, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	exeAbsPath, err := filepath.Abs(exePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	dstDir := filepath.Join(filepath.Dir(exeAbsPath), time.Now().Format("20060102-150405"))

	os.RemoveAll(dstDir)
	if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}

	// 创建目录文件, 并写入解析的目录结构
	menuDoc, err := os.Create(filepath.Join(dstDir, "menudoc.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer menuDoc.Close()

	for _, level1 := range mn.Level1List {
		fmt.Fprintln(menuDoc, strings.TrimSpace(level1.H5.Text))
		for _, level2 := range level1.Div.Level2List {
			text := strings.TrimSpace(level2.Anchor.Text)
			href := strings.TrimSpace(level2.Anchor.Href)
			if strings.HasPrefix(href, "../") {
				href = "http://mp.weixin.qq.com/wiki" + href[2:]
			}
			fmt.Fprintln(menuDoc, "\t"+text+"("+href+")")
		}
	}

	// 下载二级目录html文件
	for _, level1 := range mn.Level1List {
		for _, level2 := range level1.Div.Level2List {
			text := strings.TrimSpace(level2.Anchor.Text)
			href := strings.TrimSpace(level2.Anchor.Href)
			if strings.HasPrefix(href, "../") {
				href = "http://mp.weixin.qq.com/wiki" + href[2:]
			}

			time.Sleep(time.Second)

			if err := htmlDownload(dstDir, text, href); err != nil {
				fmt.Println("download failed", text, href)
			}
		}
	}
}

var src = []byte(`
<?xml version="1.0" encoding="utf-8"?>

<xml> 
  <!-- 新手开发者指南 -->  
  <div class="portal" id="p-.E6.96.B0.E6.89.8B.E5.BC.80.E5.8F.91.E8.80.85.E6.8C.87.E5.8D.97"> 
    <h5>
      <span class="portal_arrow"/>新手开发者指南
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E6.8E.A5.E5.85.A5.E6.8C.87.E5.8D.97">
          <a href="../17/2d4265491f12608cd170a95559800f2d.html">接入指南</a>
        </li>  
        <li id="n-.E5.85.B8.E5.9E.8B.E6.A1.88.E4.BE.8B.E4.BB.8B.E7.BB.8D">
          <a href="../7/38253a54da5b96126abea6925321f64b.html">典型案例介绍</a>
        </li>  
        <li id="n-.E5.BC.80.E5.8F.91.E8.80.85.E8.A7.84.E8.8C.83">
          <a href="../3/8b07e4a79cef674d4bcb788e1280c1b7.html">开发者规范</a>
        </li>  
        <li id="n-.E5.85.AC.E4.BC.97.E5.8F.B7.E7.B1.BB.E5.9E.8B.E7.9A.84.E6.8E.A5.E5.8F.A3.E6.9D.83.E9.99.90.E8.AF.B4.E6.98.8E">
          <a href="../7/2d301d4b757dedc333b9a9854b457b47.html">公众号类型的接口权限说明</a>
        </li>  
        <li id="n-.E5.BE.AE.E4.BF.A1.E5.BC.80.E5.8F.91.E8.80.85.E4.BA.92.E5.8A.A9.E9.97.AE.E7.AD.94.E7.B3.BB.E7.BB.9F">
          <a href="http://mp.weixin.qq.com/qa">微信开发者互助问答系统</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /新手开发者指南 -->  
  <!-- 在线调试与测试号申请 -->  
  <div class="portal" id="p-.E5.9C.A8.E7.BA.BF.E8.B0.83.E8.AF.95.E4.B8.8E.E6.B5.8B.E8.AF.95.E5.8F.B7.E7.94.B3.E8.AF.B7"> 
    <h5>
      <span class="portal_arrow"/>在线调试与测试号申请
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E6.8E.A5.E5.8F.A3.E8.B0.83.E8.AF.95.E5.B7.A5.E5.85.B7">
          <a href="http://mp.weixin.qq.com/debug/">接口调试工具</a>
        </li>  
        <li id="n-.E6.8E.A5.E5.8F.A3.E4.BD.93.E9.AA.8C.E6.B5.8B.E8.AF.95.E5.8F.B7.E7.94.B3.E8.AF.B7">
          <a href="http://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login">接口体验测试号申请</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /在线调试与测试号申请 -->  
  <!-- 返回码与报警排查 -->  
  <div class="portal" id="p-.E8.BF.94.E5.9B.9E.E7.A0.81.E4.B8.8E.E6.8A.A5.E8.AD.A6.E6.8E.92.E6.9F.A5"> 
    <h5>
      <span class="portal_arrow"/>返回码与报警排查
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E6.8E.A5.E5.8F.A3.E9.A2.91.E7.8E.87.E9.99.90.E5.88.B6.E8.AF.B4.E6.98.8E">
          <a href="../0/2e2239fa5f49388d5b5136ecc8e0e440.html">接口频率限制说明</a>
        </li>  
        <li id="n-.E5.85.A8.E5.B1.80.E6.8E.A5.E5.8F.A3.E8.BF.94.E5.9B.9E.E7.A0.81.E8.AF.B4.E6.98.8E">
          <a href="../17/fa4e1434e57290788bde25603fa2fcbd.html">全局接口返回码说明</a>
        </li>  
        <li id="n-.E6.8A.A5.E8.AD.A6.E6.8E.92.E6.9F.A5.E6.8C.87.E5.BC.95">
          <a href="../6/01405db0092f76bb96b12a9f954cd866.html">报警排查指引</a>
        </li>  
        <li id="n-.E5.BE.AE.E4.BF.A1.E6.8E.A8.E9.80.81.E6.B6.88.E6.81.AF.E4.B8.8E.E4.BA.8B.E4.BB.B6.E8.AF.B4.E6.98.8E">
          <a href="../10/b6103072d2ff5fed5fea203bcd0af256.html">微信推送消息与事件说明</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /返回码与报警排查 -->  
  <!-- 消息体签名与加解密 -->  
  <div class="portal" id="p-.E6.B6.88.E6.81.AF.E4.BD.93.E7.AD.BE.E5.90.8D.E4.B8.8E.E5.8A.A0.E8.A7.A3.E5.AF.86"> 
    <h5>
      <span class="portal_arrow"/>消息体签名与加解密
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E6.96.B9.E6.A1.88.E6.A6.82.E8.BF.B0">
          <a href="../13/80a1a25adbc46faf2716774c423b3151.html">方案概述</a>
        </li>  
        <li id="n-.E6.8E.A5.E5.85.A5.E6.8C.87.E5.BC.95">
          <a href="../0/61c3a8b9d50ac74f18bdf2e54ddfc4e0.html">接入指引</a>
        </li>  
        <li id="n-.E6.8A.80.E6.9C.AF.E6.96.B9.E6.A1.88">
          <a href="../2/3478f69c0d0bbe8deb48d66a3111ff6e.html">技术方案</a>
        </li>  
        <li id="n-.E5.BC.80.E5.8F.91.E8.80.85FAQ">
          <a href="../1/66bc8fb495f80faeacc65cfb8d931acd.html">开发者FAQ</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /消息体签名与加解密 -->  
  <!-- 基础接口 -->  
  <div class="portal" id="p-.E5.9F.BA.E7.A1.80.E6.8E.A5.E5.8F.A3"> 
    <h5>
      <span class="portal_arrow"/>基础接口
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E8.8E.B7.E5.8F.96access_token">
          <a href="../11/0e4b294685f817b95cbed85ba5e82b8f.html">获取access_token</a>
        </li>  
        <li id="n-.E8.8E.B7.E5.8F.96.E5.BE.AE.E4.BF.A1.E6.9C.8D.E5.8A.A1.E5.99.A8IP.E5.9C.B0.E5.9D.80">
          <a href="../0/2ad4b6bfd29f30f71d39616c2a0fcedc.html">获取微信服务器IP地址</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /基础接口 -->  
  <!-- 接收消息 -->  
  <div class="portal" id="p-.E6.8E.A5.E6.94.B6.E6.B6.88.E6.81.AF"> 
    <h5>
      <span class="portal_arrow"/>接收消息
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E9.AA.8C.E8.AF.81.E6.B6.88.E6.81.AF.E7.9C.9F.E5.AE.9E.E6.80.A7">
          <a href="../4/2ccadaef44fe1e4b0322355c2312bfa8.html">验证消息真实性</a>
        </li>  
        <li id="n-.E6.8E.A5.E6.94.B6.E6.99.AE.E9.80.9A.E6.B6.88.E6.81.AF">
          <a href="../10/79502792eef98d6e0c6e1739da387346.html">接收普通消息</a>
        </li>  
        <li id="n-.E6.8E.A5.E6.94.B6.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81">
          <a href="../2/5baf56ce4947d35003b86a9805634b1e.html">接收事件推送</a>
        </li>  
        <li id="n-.E6.8E.A5.E6.94.B6.E8.AF.AD.E9.9F.B3.E8.AF.86.E5.88.AB.E7.BB.93.E6.9E.9C">
          <a href="../2/f2bef3230362d18851ee22953abfadde.html">接收语音识别结果</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /接收消息 -->  
  <!-- 发送消息 -->  
  <div class="portal" id="p-.E5.8F.91.E9.80.81.E6.B6.88.E6.81.AF"> 
    <h5>
      <span class="portal_arrow"/>发送消息
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E8.A2.AB.E5.8A.A8.E5.9B.9E.E5.A4.8D.E7.94.A8.E6.88.B7.E6.B6.88.E6.81.AF">
          <a href="../14/89b871b5466b19b3efa4ada8e577d45e.html">被动回复用户消息</a>
        </li>  
        <li id="n-.E5.AE.A2.E6.9C.8D.E6.8E.A5.E5.8F.A3">
          <a href="../1/70a29afed17f56d537c833f89be979c9.html">客服接口</a>
        </li>  
        <li id="n-.E9.AB.98.E7.BA.A7.E7.BE.A4.E5.8F.91.E6.8E.A5.E5.8F.A3">
          <a href="../15/5380a4e6f02f2ffdc7981a8ed7a40753.html">高级群发接口</a>
        </li>  
        <li id="n-.E6.A8.A1.E6.9D.BF.E6.B6.88.E6.81.AF.E6.8E.A5.E5.8F.A3">
          <a href="../17/304c1885ea66dbedf7dc170d84999a9d.html">模板消息接口</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /发送消息 -->  
  <!-- 素材管理 -->  
  <div class="portal" id="p-.E7.B4.A0.E6.9D.90.E7.AE.A1.E7.90.86"> 
    <h5>
      <span class="portal_arrow"/>素材管理
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E6.96.B0.E5.A2.9E.E4.B8.B4.E6.97.B6.E7.B4.A0.E6.9D.90">
          <a href="../5/963fc70b80dc75483a271298a76a8d59.html">新增临时素材</a>
        </li>  
        <li id="n-.E8.8E.B7.E5.8F.96.E4.B8.B4.E6.97.B6.E7.B4.A0.E6.9D.90">
          <a href="../11/07b6b76a6b6e8848e855a435d5e34a5f.html">获取临时素材</a>
        </li>  
        <li id="n-.E6.96.B0.E5.A2.9E.E6.B0.B8.E4.B9.85.E7.B4.A0.E6.9D.90">
          <a href="../14/7e6c03263063f4813141c3e17dd4350a.html">新增永久素材</a>
        </li>  
        <li id="n-.E8.8E.B7.E5.8F.96.E6.B0.B8.E4.B9.85.E7.B4.A0.E6.9D.90">
          <a href="../4/b3546879f07623cb30df9ca0e420a5d0.html">获取永久素材</a>
        </li>  
        <li id="n-.E5.88.A0.E9.99.A4.E6.B0.B8.E4.B9.85.E7.B4.A0.E6.9D.90">
          <a href="../5/e66f61c303db51a6c0f90f46b15af5f5.html">删除永久素材</a>
        </li>  
        <li id="n-.E4.BF.AE.E6.94.B9.E6.B0.B8.E4.B9.85.E5.9B.BE.E6.96.87.E7.B4.A0.E6.9D.90">
          <a href="../4/19a59cba020d506e767360ca1be29450.html">修改永久图文素材</a>
        </li>  
        <li id="n-.E8.8E.B7.E5.8F.96.E7.B4.A0.E6.9D.90.E6.80.BB.E6.95.B0">
          <a href="../16/8cc64f8c189674b421bee3ed403993b8.html">获取素材总数</a>
        </li>  
        <li id="n-.E8.8E.B7.E5.8F.96.E7.B4.A0.E6.9D.90.E5.88.97.E8.A1.A8">
          <a href="../12/2108cd7aafff7f388f41f37efa710204.html">获取素材列表</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /素材管理 -->  
  <!-- 用户管理 -->  
  <div class="portal" id="p-.E7.94.A8.E6.88.B7.E7.AE.A1.E7.90.86"> 
    <h5>
      <span class="portal_arrow"/>用户管理
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E7.94.A8.E6.88.B7.E5.88.86.E7.BB.84.E7.AE.A1.E7.90.86">
          <a href="../0/56d992c605a97245eb7e617854b169fc.html">用户分组管理</a>
        </li>  
        <li id="n-.E8.AE.BE.E7.BD.AE.E7.94.A8.E6.88.B7.E5.A4.87.E6.B3.A8.E5.90.8D">
          <a href="../1/4a566d20d67def0b3c1afc55121d2419.html">设置用户备注名</a>
        </li>  
        <li id="n-.E8.8E.B7.E5.8F.96.E7.94.A8.E6.88.B7.E5.9F.BA.E6.9C.AC.E4.BF.A1.E6.81.AF.28UnionID.E6.9C.BA.E5.88.B6.29">
          <a href="../14/bb5031008f1494a59c6f71fa0f319c66.html">获取用户基本信息(UnionID机制)</a>
        </li>  
        <li id="n-.E8.8E.B7.E5.8F.96.E7.94.A8.E6.88.B7.E5.88.97.E8.A1.A8">
          <a href="../0/d0e07720fc711c02a3eab6ec33054804.html">获取用户列表</a>
        </li>  
        <li id="n-.E8.8E.B7.E5.8F.96.E7.94.A8.E6.88.B7.E5.9C.B0.E7.90.86.E4.BD.8D.E7.BD.AE">
          <a href="../8/1b86529d05db9f960e48c3c7ca5be288.html">获取用户地理位置</a>
        </li>  
        <li id="n-.E7.BD.91.E9.A1.B5.E6.8E.88.E6.9D.83.E8.8E.B7.E5.8F.96.E7.94.A8.E6.88.B7.E5.9F.BA.E6.9C.AC.E4.BF.A1.E6.81.AF">
          <a href="../17/c0f37d5704f0b64713d5d2c37b468d75.html">网页授权获取用户基本信息</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /用户管理 -->  
  <!-- 自定义菜单管理 -->  
  <div class="portal" id="p-.E8.87.AA.E5.AE.9A.E4.B9.89.E8.8F.9C.E5.8D.95.E7.AE.A1.E7.90.86"> 
    <h5>
      <span class="portal_arrow"/>自定义菜单管理
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E8.87.AA.E5.AE.9A.E4.B9.89.E8.8F.9C.E5.8D.95.E5.88.9B.E5.BB.BA.E6.8E.A5.E5.8F.A3">
          <a href="../13/43de8269be54a0a6f64413e4dfa94f39.html">自定义菜单创建接口</a>
        </li>  
        <li id="n-.E8.87.AA.E5.AE.9A.E4.B9.89.E8.8F.9C.E5.8D.95.E6.9F.A5.E8.AF.A2.E6.8E.A5.E5.8F.A3">
          <a href="../16/ff9b7b85220e1396ffa16794a9d95adc.html">自定义菜单查询接口</a>
        </li>  
        <li id="n-.E8.87.AA.E5.AE.9A.E4.B9.89.E8.8F.9C.E5.8D.95.E5.88.A0.E9.99.A4.E6.8E.A5.E5.8F.A3">
          <a href="../16/8ed41ba931e4845844ad6d1eeb8060c8.html">自定义菜单删除接口</a>
        </li>  
        <li id="n-.E8.87.AA.E5.AE.9A.E4.B9.89.E8.8F.9C.E5.8D.95.E4.BA.8B.E4.BB.B6.E6.8E.A8.E9.80.81">
          <a href="../9/981d772286d10d153a3dc4286c1ee5b5.html">自定义菜单事件推送</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /自定义菜单管理 -->  
  <!-- 帐号管理 -->  
  <div class="portal" id="p-.E5.B8.90.E5.8F.B7.E7.AE.A1.E7.90.86"> 
    <h5>
      <span class="portal_arrow"/>帐号管理
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E7.94.9F.E6.88.90.E5.B8.A6.E5.8F.82.E6.95.B0.E7.9A.84.E4.BA.8C.E7.BB.B4.E7.A0.81">
          <a href="../18/28fc21e7ed87bec960651f0ce873ef8a.html">生成带参数的二维码</a>
        </li>  
        <li id="n-.E9.95.BF.E9.93.BE.E6.8E.A5.E8.BD.AC.E7.9F.AD.E9.93.BE.E6.8E.A5.E6.8E.A5.E5.8F.A3">
          <a href="../10/165c9b15eddcfbd8699ac12b0bd89ae6.html">长链接转短链接接口</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /帐号管理 -->  
  <!-- 数据统计接口 -->  
  <div class="portal" id="p-.E6.95.B0.E6.8D.AE.E7.BB.9F.E8.AE.A1.E6.8E.A5.E5.8F.A3"> 
    <h5>
      <span class="portal_arrow"/>数据统计接口
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E7.94.A8.E6.88.B7.E5.88.86.E6.9E.90.E6.95.B0.E6.8D.AE.E6.8E.A5.E5.8F.A3">
          <a href="../3/ecfed6e1a0a03b5f35e5efac98e864b7.html">用户分析数据接口</a>
        </li>  
        <li id="n-.E5.9B.BE.E6.96.87.E5.88.86.E6.9E.90.E6.95.B0.E6.8D.AE.E6.8E.A5.E5.8F.A3">
          <a href="../8/c0453610fb5131d1fcb17b4e87c82050.html">图文分析数据接口</a>
        </li>  
        <li id="n-.E6.B6.88.E6.81.AF.E5.88.86.E6.9E.90.E6.95.B0.E6.8D.AE.E6.8E.A5.E5.8F.A3">
          <a href="../12/32d42ad542f2e4fc8a8aa60e1bce9838.html">消息分析数据接口</a>
        </li>  
        <li id="n-.E6.8E.A5.E5.8F.A3.E5.88.86.E6.9E.90.E6.95.B0.E6.8D.AE.E6.8E.A5.E5.8F.A3">
          <a href="../8/30ed81ae38cf4f977194bf1a5db73668.html">接口分析数据接口</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /数据统计接口 -->  
  <!-- 微信JS-SDK -->  
  <div class="portal" id="p-.E5.BE.AE.E4.BF.A1JS-SDK"> 
    <h5>
      <span class="portal_arrow"/>微信JS-SDK
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E5.BE.AE.E4.BF.A1JS-SDK.E8.AF.B4.E6.98.8E.E6.96.87.E6.A1.A3">
          <a href="../7/aaa137b55fb2e0456bf8dd9148dd613f.html">微信JS-SDK说明文档</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /微信JS-SDK -->  
  <!-- 微信小店接口 -->  
  <div class="portal" id="p-.E5.BE.AE.E4.BF.A1.E5.B0.8F.E5.BA.97.E6.8E.A5.E5.8F.A3"> 
    <h5>
      <span class="portal_arrow"/>微信小店接口
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E5.BE.AE.E4.BF.A1.E5.B0.8F.E5.BA.97.E6.8E.A5.E5.8F.A3">
          <a href="../8/703923b7349a607f13fb3100163837f0.html">微信小店接口</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /微信小店接口 -->  
  <!-- 微信卡券接口 -->  
  <div class="portal" id="p-.E5.BE.AE.E4.BF.A1.E5.8D.A1.E5.88.B8.E6.8E.A5.E5.8F.A3"> 
    <h5>
      <span class="portal_arrow"/>微信卡券接口
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E5.BE.AE.E4.BF.A1.E5.8D.A1.E5.88.B8.E6.8E.A5.E5.8F.A3">
          <a href="../9/d8a5f3b102915f30516d79b44fe665ed.html">微信卡券接口</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /微信卡券接口 -->  
  <!-- 微信智能接口 -->  
  <div class="portal" id="p-.E5.BE.AE.E4.BF.A1.E6.99.BA.E8.83.BD.E6.8E.A5.E5.8F.A3"> 
    <h5>
      <span class="portal_arrow"/>微信智能接口
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E8.AF.AD.E4.B9.89.E7.90.86.E8.A7.A3.E6.8E.A5.E5.8F.A3">
          <a href="../0/0ce78b3c9524811fee34aba3e33f3448.html">语义理解接口</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /微信智能接口 -->  
  <!-- 设备功能介绍 -->  
  <div class="portal" id="p-.E8.AE.BE.E5.A4.87.E5.8A.9F.E8.83.BD.E4.BB.8B.E7.BB.8D"> 
    <h5>
      <span class="portal_arrow"/>设备功能介绍
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E8.AE.BE.E5.A4.87.E5.8A.9F.E8.83.BD.E4.BB.8B.E7.BB.8D">
          <a href="../5/131b418c04b1f4fc1752f7652b14b235.html">设备功能介绍</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /设备功能介绍 -->  
  <!-- 多客服功能 -->  
  <div class="portal" id="p-.E5.A4.9A.E5.AE.A2.E6.9C.8D.E5.8A.9F.E8.83.BD"> 
    <h5>
      <span class="portal_arrow"/>多客服功能
    </h5>  
    <div class="body"> 
      <ul> 
        <li id="n-.E5.B0.86.E6.B6.88.E6.81.AF.E8.BD.AC.E5.8F.91.E5.88.B0.E5.A4.9A.E5.AE.A2.E6.9C.8D">
          <a href="../5/ae230189c9bd07a6b221f48619aeef35.html">将消息转发到多客服</a>
        </li>  
        <li id="n-.E5.AE.A2.E6.9C.8D.E7.AE.A1.E7.90.86">
          <a href="../9/6fff6f191ef92c126b043ada035cc935.html">客服管理</a>
        </li>  
        <li id="n-.E5.A4.9A.E5.AE.A2.E6.9C.8D.E4.BC.9A.E8.AF.9D.E6.8E.A7.E5.88.B6">
          <a href="../2/6c20f3e323bdf5986cfcb33cbd3b829a.html">多客服会话控制</a>
        </li>  
        <li id="n-.E8.8E.B7.E5.8F.96.E5.AE.A2.E6.9C.8D.E8.81.8A.E5.A4.A9.E8.AE.B0.E5.BD.95">
          <a href="../19/7c129ec71ddfa60923ea9334557e8b23.html">获取客服聊天记录</a>
        </li>  
        <li id="n-PC.E5.AE.A2.E6.88.B7.E7.AB.AF.E8.87.AA.E5.AE.9A.E4.B9.89.E6.8F.92.E4.BB.B6.E6.8E.A5.E5.8F.A3">
          <a href="../17/0160b650bc11ca90776343276e91082d.html">PC客户端自定义插件接口</a>
        </li> 
      </ul> 
    </div> 
  </div>  
  <!-- /多客服功能 --> 
</xml>
`)
