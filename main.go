package main

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/spf13/viper"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"test/setting"
	"time"
)

//var fs embed.FS

func main() {

	if len(os.Args) < 2 {
		fmt.Println("缺少cmd参数,程序终止")
		return
	}

	err := setting.Init(os.Args[1])
	if err != nil {
		fmt.Println("读取配置文件失败")
		return
	}

	now := time.Now()
	//玩家1
	baseMap := viper.GetString("BaseMap.url")
	team1Name := viper.GetString("Team1.Content")
	team1X := viper.GetInt("Team1.coordX")
	team1Y := viper.GetInt("Team1.coordY")
	team1Font := viper.GetInt("Team1.FontSize")
	team1FontAddress := viper.GetString("Team1.FontAddress")
	team1ImageColor := viper.GetString("Team1.Colour")

	//玩家二
	Team2Name := viper.GetString("Team2.Content")
	Team2X := viper.GetInt("Team2.coordX")
	Team2Y := viper.GetInt("Team2.coordY")
	Team2Font := viper.GetInt("Team2.FontSize")
	Team2FontAddress := viper.GetString("Team2.FontAddress")
	Team2ImageColor := viper.GetString("Team2.Colour")

	//年
	YearName := viper.GetString("Year.Content")
	YearX := viper.GetInt("Year.coordX")
	YearY := viper.GetInt("Year.coordY")
	YearFont := viper.GetInt("Year.FontSize")
	YearFontAddress := viper.GetString("Year.FontAddress")
	YearImageColor := viper.GetString("Year.Colour")

	//月
	MonthName := viper.GetString("Month.Content")
	MonthX := viper.GetInt("Month.coordX")
	MonthY := viper.GetInt("Month.coordY")
	MonthFont := viper.GetInt("Month.FontSize")
	MonthFontAddress := viper.GetString("Month.FontAddress")
	MonthImageColor := viper.GetString("Month.Colour")

	//天
	DayName := viper.GetString("Day.Content")
	DayX := viper.GetInt("Day.coordX")
	DayY := viper.GetInt("Day.coordY")
	DayFont := viper.GetInt("Day.FontSize")
	DayFontAddress := viper.GetString("Day.FontAddress")
	DayImageColor := viper.GetString("Day.Colour")

	err = poster(baseMap, team1Name, team1X, team1Y, team1Font, team1FontAddress, team1ImageColor, Team2Name, Team2X, Team2Y, Team2Font, Team2FontAddress, Team2ImageColor,
		YearName, YearX, YearY, YearFont, YearFontAddress, YearImageColor, MonthName, MonthX, MonthY, MonthFont, MonthFontAddress, MonthImageColor, DayName, DayX, DayY, DayFont, DayFontAddress, DayImageColor)
	since := time.Since(now)
	fmt.Println(since, err)

	fmt.Println("生成成功")
	time.Sleep(time.Second * 10)

}

func poster(baseMap string, team1Name string, team1X int, team1Y int, team1Font int, team1FontAddress string, team1ImageColor string, team2Name string, team2X int, team2Y int, team2Font int, team2FontAddress string, team2ImageColor string, YearName string, YearX int, YearY int, YearFont int, YearFontAddress string, YearImageColor string, MonthName string, MonthX int, MonthY int, MonthFont int, MonthFontAddress string, MonthImageColor string, DayName string, DayX int, DayY int, DayFont int, DayFontAddress string, DayImageColor string) error {

	//fmt.Println(baseMap)
	//back, err := fs.Open("./base/456.jpg")

	back, err := os.Open(baseMap)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	badec, err := jpeg.Decode(back)
	if err != nil {
		return err
	}
	rgba := image.NewRGBA(badec.Bounds())
	draw.Draw(rgba, rgba.Bounds(), badec, image.Point{
		X: 0,
		Y: 0,
	}, draw.Src)
	err = drawText(rgba, team1Name, team1X, team1Y, team1Font, team1ImageColor, team1FontAddress)
	if err != nil {
		return err
	}
	err = drawText(rgba, team2Name, team2X, team2Y, team2Font, team2ImageColor, team2FontAddress)

	err = drawText(rgba, YearName, YearX, YearY, YearFont, YearImageColor, YearFontAddress)
	err = drawText(rgba, MonthName, MonthX, MonthY, MonthFont, MonthImageColor, MonthFontAddress)
	err = drawText(rgba, DayName, DayX, DayY, DayFont, DayImageColor, DayFontAddress)

	create, err := os.Create("image/" + time.Now().Format("20060102150405") + ".png")
	if err != nil {
		return err
	}
	err = png.Encode(create, rgba)
	return err
}

func drawText(dst draw.Image, s string, x, y int, fontSize int, imageColor string, FontAddress string) error {
	file, err := ioutil.ReadFile(FontAddress)

	if err != nil {
		return err
	}
	font, err := freetype.ParseFont(file)
	if err != nil {
		return err
	}
	con := freetype.NewContext()
	con.SetFont(font)
	con.SetFontSize(float64(fontSize))
	con.SetDst(dst)
	if imageColor == "white" {
		con.SetSrc(image.White)
	} else {
		con.SetSrc(image.Black)

	}
	con.SetClip(dst.Bounds())
	_, err = con.DrawString(s, freetype.Pt(x, y))
	return err
}
