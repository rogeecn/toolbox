package urlx

import "net/url"

type URL struct {
	u     *url.URL
	query url.Values
}

func New(u *url.URL) *URL {
	return &URL{u: u}
}

// FromString
func FromString(s string) (*URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	return New(u), nil
}

func (u *URL) String() string {
	if u.query != nil {
		u.u.RawQuery = u.query.Encode()
	}
	return u.u.String()
}

func (u *URL) Path() string {
	return u.u.Path
}

func (u *URL) Query() url.Values {
	if u.query == nil {
		u.query = u.u.Query()
	}
	return u.query
}

func (u *URL) RawQuery() string {
	return u.u.RawQuery
}

func (u *URL) GetQuery(key string) string {
	return u.u.Query().Get(key)
}

func (u *URL) Fragment() string {
	return u.u.Fragment
}

func (u *URL) Scheme() string {
	return u.u.Scheme
}

func (u *URL) Host() string {
	return u.u.Host
}

func (u *URL) User() string {
	return u.u.User.String()
}

func (u *URL) Username() string {
	return u.u.User.Username()
}

func (u *URL) Password() string {
	p, _ := u.u.User.Password()
	return p
}

// SetPath
func (u *URL) SetPath(path string) *URL {
	u.u.Path = path
	return u
}

// SetQuery
func (u *URL) SetQuery(query url.Values) *URL {
	u.u.RawQuery = query.Encode()
	return u
}

// SetRawQuery
func (u *URL) SetRawQuery(rawQuery string) *URL {
	u.u.RawQuery = rawQuery
	return u
}

// SetQueryValue
func (u *URL) SetQueryValue(key, value string) *URL {
	u.query.Set(key, value)
	return u
}

func (u *URL) UnsetQueryValue(key string) *URL {
	u.query.Del(key)
	return u
}

// SetFragment
func (u *URL) SetFragment(fragment string) *URL {
	u.u.Fragment = fragment
	return u
}

// SetScheme
func (u *URL) SetScheme(scheme string) *URL {
	u.u.Scheme = scheme
	return u
}

// SetHost
func (u *URL) SetHost(host string) *URL {
	u.u.Host = host
	return u
}

// SetUser
func (u *URL) SetUser(user string) *URL {
	u.u.User = url.User(user)
	return u
}

// SetUsername
func (u *URL) SetUsername(username string) *URL {
	u.u.User = url.User(username)
	return u
}

// SetPassword
func (u *URL) SetPassword(password string) *URL {
	u.u.User = url.UserPassword(u.Username(), password)
	return u
}

// SetUsernamePassword
func (u *URL) SetUsernamePassword(username, password string) *URL {
	u.u.User = url.UserPassword(username, password)
	return u
}

// SetUserinfo
func (u *URL) SetUserinfo(userinfo *url.Userinfo) *URL {
	u.u.User = userinfo
	return u
}
