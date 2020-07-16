package hellomock

import "testing"

func TestCompany_Meeting(t *testing.T) {
	person:=NewPersobn("王你妹")
	company:=NewCompany(person)
	t.Log(company.Meeting("王尼玛"))
}
