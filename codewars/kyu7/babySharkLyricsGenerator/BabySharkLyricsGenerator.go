package kyu7

func BabySharkLyrics() string{
	c:=""
	a:=[]string{"Baby shark","Mommy shark","Daddy shark","Grandma shark","Grandpa shark","Let's go hunt",}
	for _,w:=range a{
		for b:=0;b<3;b++{c+=w+", doo doo doo doo doo doo\n"}
		c+=w+"!\n"
	}
  return c+"Run away,…\n"
}
