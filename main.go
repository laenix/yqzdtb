package main

import (
	"bytes"
	"errors"
	"net/http"
	"strings"
	"github.com/robfig/cron"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	var wg sync.WaitGroup //用于等待所有协程都完成
	wg.Add(1)
	go Autoreport()
	wg.Wait()
}
func Autoreport() {
	c := cron.New()
	c.AddFunc("* 01 8 * * *", Report) //每天八点零一自动填报
	c.Start()
}
func Report() {
	cookie := Getcookie("学号", "身份证号码后六位") //这里更换成你的学号和密码
	postdata := Getinfo(cookie)
	Postdata(cookie, postdata)
}
func noRedirect(req *http.Request, via []*http.Request) error {
	return errors.New("Don't redirect!")
}

func Getcookie(username string, password string) string {
	url := "http://xg.sylu.edu.cn/SPCP/Web"
	html, _ := goquery.NewDocument(url)
	findhidv := html.Find("input[name=ReSubmiteFlag]")
	hidv, _ := findhidv.Attr("value")
	findcode := html.Find("#code-box")
	findcode.SetText("1234")
	Parmers := "ReSubmiteFlag=" + hidv + "&StuLoginMode=1&txtUid=" + username + "&txtPwd=" + password + "&codeInput=1234"
	postdata := []byte(Parmers)
	client := &http.Client{
		CheckRedirect: noRedirect,
	}
	client.Head("User-Agent Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.8 Safari/537.36")
	resp, _ := client.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer(postdata))
	defer resp.Body.Close()
	cookies := resp.Header.Values("Set-Cookie")
	cookie1 := cookies[0][0:43]
	cookie2 := cookies[1][0:366]
	return cookie1 + " " + cookie2
}

func Getinfo(cookie string) string {
	url := "http://xg.sylu.edu.cn/SPCP/Web/Report/Index"
	client := &http.Client{}
	var req *http.Request
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", cookie)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	html, _ := goquery.NewDocumentFromReader(resp.Body)
	//学号
	findStudentId := html.Find("input[name=StudentId]")
	StudentId, _ := findStudentId.Attr("value")
	//姓名
	findName := html.Find("input[name=Name]")
	Name, _ := findName.Attr("value")
	//性别
	findSex := html.Find("input[name=Sex]")
	Sex, _ := findSex.Attr("value")
	//学院
	findCollegeNo := html.Find("input[name=CollegeNo]")
	CollegeNo, _ := findCollegeNo.Attr("value")
	//年级
	findSpeGrade := html.Find("input[name=SpeGrade]")
	SpeGrade, _ := findSpeGrade.Attr("value")
	//专业名称
	findSpecialtyName := html.Find("input[name=SpecialtyName]")
	SpecialtyName, _ := findSpecialtyName.Attr("value")
	//班级名称
	findClassName := html.Find("input[name=ClassName]")
	ClassName, _ := findClassName.Attr("value")
	//手机号
	findMoveTel := html.Find("input[name=MoveTel]")
	MoveTel, _ := findMoveTel.Attr("value")
	/*
		//当前所在地址 省
		findProvince := html.Find("select[name=Province]")
		Province, _ := findProvince.Attr("value")
		//当前所在地址 市
		findCity := html.Find("select[name=City]")
		City, _ := findCity.Attr("value")
		//当前所在地址 县
		findCounty := html.Find("select[name=County]")
		County, _ := findCounty.Attr("value")
	*/
	//当前所在地址 哪
	findComeWhere := html.Find("input[name=ComeWhere]")
	ComeWhere, _ := findComeWhere.Attr("value")
	/*
		//家庭住址： 省
		findFaProvince := html.Find("input[name=FaProvince]")
		FaProvince, _ := findFaProvince.Attr("value")
		//家庭住址： 市
		findFaCity := html.Find("input[name=FaCity]")
		FaCity, _ := findFaCity.Attr("value")
		//家庭住址： 县
		findFaCounty := html.Find("input[name=FaCounty]")
		FaCounty, _ := findFaCounty.Attr("value")
	*/
	//家庭住址： 哪
	findFaComeWhere := html.Find("input[name=FaComeWhere]")
	FaComeWhere, _ := findFaComeWhere.Attr("value")
	//身份证号
	findIdCard := html.Find("input[name=IdCard]")
	IdCard, _ := findIdCard.Attr("value")

	//家庭所在地 省
	findFaProvinceName := html.Find("input[name=FaProvinceName]")
	FaProvinceName, _ := findFaProvinceName.Attr("value")
	//家庭所在地 市
	findFaCityName := html.Find("input[name=FaCityName]")
	FaCityName, _ := findFaCityName.Attr("value")
	//家庭所在地 县
	findFaCountyName := html.Find("input[name=FaCountyName]")
	FaCountyName, _ := findFaCountyName.Attr("value")

	//radio
	radio1 := "23fafd17-1db3-4801-9332-7714253a9582"
	radio2 := "562d72d2-29f6-4769-a171-e3355a0645e3"
	radio3 := "59e3a901-61d3-458d-9df3-355d38580873"
	radio4 := "18ea1afc-e743-436a-9cb4-67ccbd106b65"
	radio5 := "48df00db-1f83-4ba4-837d-9aa5a177a8a6"
	//hidv
	findhidv := html.Find("input[name=ReSubmiteFlag]")
	hidv, _ := findhidv.Attr("value")
	PZData := `[
		{"OptionName":"否","SelectId":"23fafd17-1db3-4801-9332-7714253a9582","TitleId":"6588aa92-9d42-44fb-a159-2b13d55ef605","OptionType":"0"},
		{"OptionName":"否","SelectId":"562d72d2-29f6-4769-a171-e3355a0645e3","TitleId":"7734dd93-de9e-4482-9b4c-302de877d205","OptionType":"0"},
		{"OptionName":"否","SelectId":"59e3a901-61d3-458d-9df3-355d38580873","TitleId":"d55cae21-6a3b-4816-b5ea-26a1e35f366a","OptionType":"0"},
		{"OptionName":"否","SelectId":"18ea1afc-e743-436a-9cb4-67ccbd106b65","TitleId":"38877412-d711-494a-a136-7b2e85a93efd","OptionType":"0"},
		{"OptionName":"否","SelectId":"48df00db-1f83-4ba4-837d-9aa5a177a8a6","TitleId":"fc02846e-8bb7-4436-a9ce-edf822d82680","OptionType":"0"},
		{"OptionName":"36.3","SelectId":"","TitleId":"f90f409a-f6ba-4707-81a1-24e276a5bcc7","OptionType":"2"}
	]`
	postdata := "" +
		"StudentId=" + StudentId +
		"&Name=" + Name +
		"&Sex=" + Sex +
		"&SpeType=B&CollegeNo=" + CollegeNo +
		"&SpeGrade=" + SpeGrade +
		"&SpecialtyName=" + SpecialtyName +
		"&ClassName=" + ClassName +
		"&MoveTel=" + MoveTel +
		"&Province=" + IdCard[0:2] + "0000" +
		"&City=" + IdCard[0:4] + "00" +
		"&County=" + IdCard[0:6] +
		"&ComeWhere=" + ComeWhere +
		"&FaProvince=" + IdCard[0:2] + "0000" +
		"&FaCity=" + IdCard[0:4] + "00" +
		"&FaCounty=" + IdCard[0:6] +
		"&FaComeWhere=" + FaComeWhere +
		"&radio_1=" + radio1 +
		"&radio_2=" + radio2 +
		"&radio_3=" + radio3 +
		"&text_1=" + "36.3" +
		"&radio_4=" + radio4 +
		"&radio_5=" + radio5 +
		"&Other=&GetAreaUrl=/SPCP/Web/Report/GetArea" +
		"&IdCard=" + IdCard +
		"&ProvinceName=" + FaProvinceName +
		"&CityName=" + FaCityName +
		"&CountyName=" + FaCountyName +
		"&FaProvinceName=" + FaProvinceName +
		"&FaCityName=" + FaCityName +
		"&FaCountyName=" + FaCountyName +
		"&radioCount=5&checkboxCount=0&blackCount=1&PZData=" + PZData +
		"&ReSubmiteFlag=" + hidv
	return postdata
}

func Postdata(cookie string, postdata string) {
	client := &http.Client{}
	url := "http://xg.sylu.edu.cn/SPCP/Web/Report/Index"
	req, _ := http.NewRequest("POST", url, strings.NewReader(postdata))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", cookie)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}
