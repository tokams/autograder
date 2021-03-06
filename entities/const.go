package git

const (
	// Standard repository names
	StandardRepoName   = "labs"
	GroupsRepoName     = "glabs"
	CourseInfoName     = "course-info"
	TestRepoName       = "labs-test"
	GrouptestRepoName  = "groups-test" // Deprecated. Use only the TEST_REPO_NAME.
	CodeReviewRepoName = "code-reviews"
	GroupRepoPrefix    = "group"

	// Standard team names
	StudentsTeam = "students"
	OwnersTeam   = "Owners"

	// Team premission names on github
	AdminPermission = "admin"
	PullPermission  = "pull"
	PushPermission  = "push"
)

const (
	// Assignment types
	IndividualType int = iota
	GroupType
)

// IgnoreFileContent is the standard content for the autogenerated .gitignore file.
const IgnoreFileContent = "# Compiled source #\n" +
	"###################\n" +
	"*.com\n" +
	"*.class\n" +
	"*.dll\n" +
	"*.exe\n" +
	"*.o\n" +
	"*.so\n" +
	"**/main\n" +
	"\n" +
	"# Packages #\n" +
	"############\n" +
	"# it's better to unpack these files and commit the raw source\n" +
	"# git has its own built in compression methods\n" +
	"*.7z\n" +
	"*.dmg\n" +
	"*.gz\n" +
	"*.iso\n" +
	"*.jar\n" +
	"*.rar\n" +
	"*.tar\n" +
	"*.zip\n" +
	"\n" +
	"# Logs and databases #\n" +
	"######################\n" +
	"*.log\n" +
	"*.sql\n" +
	"*.sqlite\n" +
	"\n" +
	"# OS generated files #\n" +
	"######################\n" +
	".DS_Store\n" +
	".DS_Store?\n" +
	"._*\n" +
	".Spotlight-V100\n" +
	".Trashes\n" +
	"ehthumbs.db\n" +
	"Thumbs.db\n" +
	"\n" +
	"# Text Editors #\n" +
	"################\n" +
	"**/*~\n" +
	".project\n" +
	"*#\n" +
	"**/*.bak\n" +
	".classpath\n" +
	".settings\n" +
	".idea\n" +
	".metadata\n" +
	"*.iml\n" +
	"*.ipr\n" +
	"*.swp\n" +
	"*.swo\n"
