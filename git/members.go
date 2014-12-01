package git

import (
	"encoding/gob"
	"errors"
	"log"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	"github.com/hfurubotten/diskv"
)

func init() {
	gob.Register(Member{})
}

var userstore = diskv.New(diskv.Options{
	BasePath:     "diskv/users/",
	CacheSizeMax: 1024 * 1024 * 256,
})

type Member struct {
	githubclient *github.Client
	Username     string
	Name         string
	StudentID    int
	IsTeacher    bool
	IsAdmin      bool

	Teaching         []string
	Courses          []string
	AssistantCourses []string

	accessToken token
}

func NewMember(oauthtoken string) (m Member) {
	m = Member{accessToken: NewToken(oauthtoken)}

	var err error
	if m.accessToken.HasTokenInStore() {
		m.Username, err = m.accessToken.GetUsernameFromTokenInStore()
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		err = m.loadDataFromGithub()
		if err != nil {
			log.Println(err)
			return
		}
	}

	err = m.loadData()
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func NewMemberFromUsername(username string) (m Member) {
	m = Member{}
	m.Username = username

	err := m.loadData()
	if err != nil {
		log.Println(err)
	}

	return
}

func (m *Member) loadDataFromGithub() (err error) {
	err = m.connectToGithub()
	if err != nil {
		return
	}

	user, _, err := m.githubclient.Users.Get("")
	if err != nil {
		return
	}

	if user.Login != nil {
		m.Username = *user.Login
	}

	if user.Name != nil {
		m.Name = *user.Name
	}

	return
}

func (m *Member) loadData() (err error) {
	if userstore.Has(m.Username) {
		var tmp Member

		err = userstore.ReadGob(m.Username, &tmp, false)
		if err != nil {
			return
		}
		m.Copy(tmp)

		if !m.accessToken.HasTokenInStore() {
			m.accessToken.SetUsernameToTokenInStore(m.Username)
		}
	}

	return
}

func (m Member) StickToSystem() (err error) {
	return userstore.WriteGob(m.Username, m)
}

func (m *Member) Copy(tmp Member) {
	m.Username = tmp.Username
	m.Name = tmp.Name
	m.StudentID = tmp.StudentID
	m.IsTeacher = tmp.IsTeacher
	m.IsAdmin = tmp.IsAdmin
	m.Teaching = tmp.Teaching
	m.Courses = tmp.Courses
	m.AssistantCourses = tmp.AssistantCourses
}

func (m Member) IsComplete() bool {
	if m.Name == "" || m.StudentID == 0 || m.Username == "" {
		return false
	}

	return true
}

func (m *Member) connectToGithub() error {
	if m.githubclient != nil {
		return nil
	}

	if !m.accessToken.HasToken() {
		return errors.New("Missing AccessToken to the memeber. Can't contact github.")
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: m.accessToken.GetToken()},
	}
	m.githubclient = github.NewClient(t.Client())
	return nil
}

func (m *Member) ListOrgs() (ls []string, err error) {
	err = m.connectToGithub()
	if err != nil {
		return
	}

	orgs, _, err := m.githubclient.Organizations.List("", nil)

	ls = make([]string, len(orgs))

	for i, org := range orgs {
		ls[i] = *org.Login
	}

	return
}

func (m *Member) AddOrganization(org Organization) (err error) {
	if m.Courses == nil {
		m.Courses = make([]string, 0)
	}

	m.Courses = append(m.Courses, org.Name)

	err = org.AddMembership(*m)

	return
}

func (m *Member) AddTeachingOrganization(org Organization) (err error) {
	if m.Courses == nil {
		m.Teaching = make([]string, 0)
	}

	m.IsTeacher = true
	m.Teaching = append(m.Courses, org.Name)

	return
}

func ListAllMembers() (out []Member) {
	out = make([]Member, 0)
	keys := userstore.Keys()
	var m Member

	for key := range keys {
		m = NewMemberFromUsername(key)
		out = append(out, m)
	}

	return
}