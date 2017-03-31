package openbd

import "testing"

const (
	host = "https://api.openbd.jp"
	isbn = "9784780802047"
)

func TestNewClient(t *testing.T) {
	cli, err := NewClientV1("", nil)
	if err == nil {
		t.Fatal("an error is nil")
	}
	cli, err = NewClientV1(host, nil)
	if err != nil {
		t.Fatal(err)
	}
	if cli == nil {
		t.Fatal("a client is nil")
	}
}

func TestClient_Get(t *testing.T) {
	cli, err := NewClientV1(host+"p", nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = cli.Get(isbn)
	if err == nil {
		t.Fatal("an error is nil")
	}
	cli, err = NewClientV1(host, nil)
	if err != nil {
		t.Fatal(err)
	}
	m, err := cli.Get("123")
	if err != nil{
		t.Fatal(err)
	}
	if bi, ok := m["123"]; ok {
		t.Fatalf("go %v, want nil", bi)
	}
	m, err = cli.Get(isbn)
	if err != nil {
		t.Fatal(err)
	}
	if len(m) == 0 {
		t.Fatal("a response is nil")
	}
	bi, ok := m[isbn]
	if !ok {
		t.Fatalf("not found - isbn = %s", isbn)
	}
	if bi.ISBN() != isbn {
		t.Errorf("got %s, want %s", bi.ISBN(), isbn)
	}
	if bi.Title() != "おにぎりレシピ101" {
		t.Errorf("got %s, want %s", bi.Title(), "おにぎりレシピ101")
	}
	if bi.Author() != "山田玲子／著 水野菜生／英訳" {
		t.Errorf("got %s, want %s", bi.Author(), "山田玲子／著 水野菜生／英訳")
	}
	if bi.Cover() != "https://cover.openbd.jp/9784780802047.jpg" {
		t.Errorf("got %s, want %s", bi.Cover(), "https://cover.openbd.jp/9784780802047.jpg")
	}
	if bi.PublishedAt() != "20140408" {
		t.Errorf("got %s, want %s", bi.PublishedAt(), "20140408")
	}
	if bi.Publisher() != "ポット出版" {
		t.Errorf("got %s, want %s", bi.Publisher(), "ポット出版")
	}
	if bi.Series() != "" {
		t.Errorf("got %s, want %s", bi.Series(), "")
	}
	if bi.Volume() != "" {
		t.Errorf("got %s, want %s", bi.Volume(), "")
	}
	if bi.ShortDescription() != "海外でも人気の日本のソウルフード、おにぎり。クッキングアドバイザー・山田玲子が考えた101のレシピを英訳付きでご紹介します。" {
		t.Errorf("got %s, want %s", bi.ShortDescription(), "海外でも人気の日本のソウルフード、おにぎり。クッキングアドバイザー・山田玲子が考えた101のレシピを英訳付きでご紹介します。")
	}
	if bi.Description() != "101人いれば、101通り、好みのおにぎりがあります。\nマイおにぎりを作ってもらうためのヒントになればと、クッキングアドバイザー・山田玲子が101のおにぎりレシピを考えました。全文英訳付き。\n日本のソウルフード、easy、simple、healthyなおにぎりは海外でも人気です。\n外国の方へのプレゼントなど、小さな外交がこの本から始まります。\n\nOnigiri—a healthy fast food—is the soul food of the Japanese. Although it may not be as widely recognized as sushi, onigiri is synonymous with the phrase “taste of home,” and is a staple of Japanese comfort food. Its simplicity—just combining rice and toppings—offers endless possibilities without borders. The portable onigiri can be served in all kinds of situations. It’s perfect for bento lunch, as a light snack, or even as party food. \nReiko Yamada’s 101 simple and easy riceball (onigiri) recipes include mixed, grilled, sushi-style onigiri and more! This cookbook is a perfect introduction to the art of onigiri-making, filled with unique recipes that are bound to inspire your Japanese culinary creativity. Pick up a copy, and you’ll become an onigiri expert in no time!\n\n" {
		t.Errorf("got %s, want %s", bi.Description(), "101人いれば、101通り、好みのおにぎりがあります。\nマイおにぎりを作ってもらうためのヒントになればと、クッキングアドバイザー・山田玲子が101のおにぎりレシピを考えました。全文英訳付き。\n日本のソウルフード、easy、simple、healthyなおにぎりは海外でも人気です。\n外国の方へのプレゼントなど、小さな外交がこの本から始まります。\n\nOnigiri—a healthy fast food—is the soul food of the Japanese. Although it may not be as widely recognized as sushi, onigiri is synonymous with the phrase “taste of home,” and is a staple of Japanese comfort food. Its simplicity—just combining rice and toppings—offers endless possibilities without borders. The portable onigiri can be served in all kinds of situations. It’s perfect for bento lunch, as a light snack, or even as party food. \nReiko Yamada’s 101 simple and easy riceball (onigiri) recipes include mixed, grilled, sushi-style onigiri and more! This cookbook is a perfect introduction to the art of onigiri-making, filled with unique recipes that are bound to inspire your Japanese culinary creativity. Pick up a copy, and you’ll become an onigiri expert in no time!\n\n")
	}
	if bi.JBPADescription() != "" {
		t.Errorf("got %s, want %s", bi.JBPADescription(), "")
	}
	if bi.TableOfContents() != "はじめに　INTRODUCTION ……002\nこの本の使い方　HOW TO USE THIS COOKBOOK……002\nご飯の炊き方　HOW TO COOK RICE ……008\n塩のこと　ALL ABOUT SALT ……010\nおにぎりのにぎり方　HOW TO MAKE ONIGIRI ……012\nおにぎりと具材の相性　ONIGIRI COMBINATIONS……014\nのりの使い方　HOW TO USE  NORI SEAWEED ……016\n\n1. 中に入れる FILL……019\n2. 混ぜる MIX……037\n3. 炊き込む／焼く TAKIKOMI / GRILL……067\n4. のせる／つつむ SUSHI-STYLE / WRAP……085\n\nおかずレシピ　OKAZU RECIPE……106\nおにぎりを冷凍保存　HOW TO KEEP LEFTOVER ONIGIRI ……113\nおにぎりをお弁当に詰める　HOW TO PACK A BENTO ……104\nINGREDIENT GLOSSARY ……116\nこの本で使用したもの　PRODUCT LIST ……120\nプロフィール　PROFILE……121\nおわりに　CLOSING ……122" {
		t.Errorf("got %s, want %s", bi.TableOfContents(), "はじめに　INTRODUCTION ……002\nこの本の使い方　HOW TO USE THIS COOKBOOK……002\nご飯の炊き方　HOW TO COOK RICE ……008\n塩のこと　ALL ABOUT SALT ……010\nおにぎりのにぎり方　HOW TO MAKE ONIGIRI ……012\nおにぎりと具材の相性　ONIGIRI COMBINATIONS……014\nのりの使い方　HOW TO USE  NORI SEAWEED ……016\n\n1. 中に入れる FILL……019\n2. 混ぜる MIX……037\n3. 炊き込む／焼く TAKIKOMI / GRILL……067\n4. のせる／つつむ SUSHI-STYLE / WRAP……085\n\nおかずレシピ　OKAZU RECIPE……106\nおにぎりを冷凍保存　HOW TO KEEP LEFTOVER ONIGIRI ……113\nおにぎりをお弁当に詰める　HOW TO PACK A BENTO ……104\nINGREDIENT GLOSSARY ……116\nこの本で使用したもの　PRODUCT LIST ……120\nプロフィール　PROFILE……121\nおわりに　CLOSING ……122")
	}
}
