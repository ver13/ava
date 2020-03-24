package string_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	. "github.com/ver13/ava/pkg/common/string"
)

type stringSuite struct {
	suite.Suite

	str *String
}

func TestStringInit(t *testing.T) {
	suite.Run(t, new(stringSuite))
}

func (s *stringSuite) BeforeTest() {
	s.T().Log("BeforeTest")
}

func (s *stringSuite) AfterTest() {
	s.T().Log("AfterTest")
}

func (s *stringSuite) SetupSuite() {
	s.T().Log("SetupSuite")
	s.str = GetInstance()
}

func (s *stringSuite) SetupTest() {
	s.T().Log("SetupTest")
}

func (s *stringSuite) TearDownSuite() {
	s.T().Log("TearDownSuite")
}

func (s *stringSuite) TearDownTest() {
	s.T().Log("TearDownTest")
}

/*
func (s *stringSuite) TestString_Filter() {
	Convey("Given a struct ", s.T(), func() {
		hFunc := func(s string) bool {
			return strings.HasPrefix(s, "h")
		}
		result := s.str.Filter(hFunc, []string{"cheese", "mouse", "hi", "there", "horse"})
		correct := []string{"hi", "horse"}
		So(len(result), ShouldEqual, len(correct))

		for i := range result {
			So(result[i], ShouldResemble, correct[i])
		}
	})
}

func (s *stringSuite) TestString_FindBetween() {
	Convey("Given a str ", s.T(), func() {
		str := "Hello <em>World</em>!"

		between, remainder, found := s.str.FindBetween(str, "<em>", "</em>")
		So(between, ShouldResemble, "World")
		So(remainder, ShouldResemble, "!")
		So(found, ShouldBeTrue)

		between, remainder, found = s.str.FindBetween(str, "l", "l")
		So(between, ShouldResemble, "")
		So(remainder, ShouldResemble, "o <em>World</em>!")
		So(found, ShouldBeTrue)

		between, remainder, found = s.str.FindBetween(str, "<i>", "</i>")
		So(between, ShouldResemble, "")
		So(remainder, ShouldResemble, "Hello <em>World</em>!")
		So(found, ShouldBeFalse)
	})
}

func (s *stringSuite) TestString_StripHTMLTags() {
	Convey("Given a str ", s.T(), func() {
		withHTML := "<div>Hello > World <br/> <im src='xxx'/>"
		skippedHTML := "Hello > World  "

		So(s.str.StripHTMLTags(withHTML), ShouldResemble, skippedHTML)
	})
}

func (s *stringSuite) TestString_ReplaceHTMLTags() {
	Convey("Given a str ", s.T(), func() {
		withHTML := "<div>Hello > World <br/> <im src='xxx'/>"
		replacedHTML := "xxHello > World xx xx"

		So(s.str.ReplaceHTMLTags(withHTML, "xx"), ShouldResemble, replacedHTML)
	})
}
*/

func (s *stringSuite) TestString_MarshalJSON() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ListContains() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ListContainsCaseInsensitive() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_PrettifyJSON() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_EscapeJSON() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_StripHTMLTags() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ReplaceHTMLTags() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_MD5Hex() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_SHA1Base64() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_AddURLParam() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ConvertTime() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_CSV() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ToInt() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ToFloat() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ToBool() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_InSlice() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_JoinFormat() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_Join() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_FormatBigInt() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_FormatMemory() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ReplaceMulti() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ToUpperCamelCase() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ToLowerCamelCase() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ToLower() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_ToUpper() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_MapSortedKeys() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_MapGroupedNumberPostfixSortedKeys() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_MapGroupedNumberPostfixSortedValues() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_EndsWithNumber() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_SplitNumberPostfix() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_SplitOnce() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_SplitOnceChar() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_SplitOnceRune() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_MapFunc() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_Filter() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_FindBetween() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (s *stringSuite) TestString_Find() {
	Convey("Given a ", s.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
