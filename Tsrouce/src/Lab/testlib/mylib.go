package testlib

var com map[string]string

// 개발자가 패키지를 만들때 사용하면 됨
// 패키지 실행시 처음으로 실행(호출)되는 함수
func init() { // 패키지 로드 시 초기화 할 때 사용되는 함수
	com = make(map[string]string)
	com["KTDS"] = "KTDS Inc"
	com["KT"] = "KT Inc"
	com["FB"] = "FaceBook Inc"
	com["GO"] = "Google Inc"
}

func GetCompany(name string) string { // 외부에서 접근 가능
	return com[name]
}

func myKey() { // 내부에서만 접근 가능

}
