package gssion

func (sh *Context) Login() {
	sh.IsAuthenticated = true
	sh.setCookie(sh.w, sh.r)
	store(sh)
}

func (sh *Context) Logout() {
	RemoveCookie("sessionid", sh.w, sh.r)
}

func (sh *Context) IsLoggedIn() bool {
	return sh.IsAuthenticated
}
