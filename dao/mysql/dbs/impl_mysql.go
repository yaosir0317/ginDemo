package dbs

import (
	"context"
	"fmt"
	"time"

	"github.com/gocraft/dbr/v2"
)

type CommonField struct {
	TimeCreated     time.Time `db:"time_created"`
	TimeLastUpdated time.Time `db:"time_last_updated"`
}

type UserSexInt int64
type UserSexStr string

var categoryMap = map[int64]string{
	1: "搞笑",
	2: "新奇酷炫",
	3: "影视娱乐",
	4: "游戏动漫",
	5: "互动",
	6: "内涵",
	7: "美女",
	8: "其他",
	9: "搞笑内涵",
}

const (
	Unclear    UserSexInt = 0
	Male       UserSexInt = 1
	FeMale     UserSexInt = 2
	UnclearStr UserSexStr = "未选择"
	MaleStr    UserSexStr = "男"
	FeMaleStr  UserSexStr = "女"
)

type User struct {
	UserId   string     `db:"user_id"`
	NickName string     `db:"nick_name"`
	SexNum   UserSexInt `db:"sex"`
	Sex      UserSexStr
	Mobile   string `db:"mobile"`
	CommonField
}

func (u *User) FormatUserSexStr() {
	switch u.SexNum {
	case Unclear:
		u.Sex = UnclearStr
	case Male:
		u.Sex = MaleStr
	case FeMale:
		u.Sex = FeMaleStr
	default:
		u.Sex = UnclearStr
	}
}

type PAuthorSecondCategory struct {
	Id        int64  `db:"id"`
	UserId    string `db:"user_id"`
	TagId     int64  `db:"tag_id"`
	MarkThree int64  `db:"mark_three"`
	CommonField
}

type PArticle struct {
	ArticleId string `db:"article_id"`
	Title     string `db:"title"`
	Verify    int64  `db:"verify"`
	Deleted   int64  `db:"deleted"`
	Duration  int64  `db:"duration"`
	Category  int64  `db:"category"`
	Rate      int64  `db:"rate"`
	CommonField
}

func (i *impl) GetCount(ctx context.Context) (uint64, error) {
	session := i.dbr.NewSession(nil)
	builder := session.Select("count(*)").From("p_author_second_category")
	buffer := dbr.NewBuffer()
	err := builder.Build(i.dialect, buffer)
	if err != nil {
		return 0, err
	}
	var cnt uint64
	err = builder.LoadOneContext(ctx, &cnt)
	//err = builder.LoadOneContext(ctx, cnt)
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func (i *impl) GetRow(ctx context.Context) {
	session := i.dbr.NewSession(nil)
	builder := session.Select("*").From("p_author_second_category").Limit(1)
	buffer := dbr.NewBuffer()
	err := builder.Build(i.dialect, buffer)
	if err != nil {
		fmt.Println(err)
	}
	var p *PAuthorSecondCategory
	err = builder.LoadOneContext(ctx, &p)
	fmt.Println(p)
}

func (i *impl) GetArticle(ctx context.Context, articleId string) (*PArticle, error) {
	session := i.dbr.NewSession(nil)
	builder := session.Select("*").From("p_article").Where(dbr.Eq("article_id", articleId))
	buffer := dbr.NewBuffer()
	err := builder.Build(i.dialect, buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var p *PArticle
	err = builder.LoadOneContext(ctx, &p)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return p, nil
}

func (i *impl) GetUser(ctx context.Context, userId, phone string) (*User, error) {
	session := i.dbr.NewSession(nil)
	var queries []dbr.Builder
	queries = append(queries, dbr.Eq("user_id", userId))
	queries = append(queries, dbr.Eq("mobile", phone))
	query := dbr.Or(queries...)
	builder := session.Select("*").From("p_user").Where(query)
	buffer := dbr.NewBuffer()
	err := builder.Build(i.dialect, buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var u *User
	err = builder.LoadOneContext(ctx, &u)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	u.FormatUserSexStr()
	return u, nil
}

func (i *impl) GetTitleMap(ctx context.Context, articleIds []string) (map[string]string, error) {
	session := i.dbr.NewSession(nil)
	builder := session.Select("article_id", "title").From("p_article").Where(dbr.Expr(fmt.Sprintf("`article_id` IN ?"), articleIds))
	buffer := dbr.NewBuffer()
	err := builder.Build(i.dialect, buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var ret []*PArticle
	_, err = builder.LoadContext(ctx, &ret)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := make(map[string]string, len(articleIds))
	if ret != nil {
		for _, t := range ret {
			result[t.ArticleId] = t.Title
		}
	}
	return result, nil
}

func (i *impl) GetArticleCategoryMap(ctx context.Context, articleIds []string) (map[string]string, error) {
	session := i.dbr.NewSession(nil)
	builder := session.Select("article_id", "category").From("p_article").Where(dbr.Expr(fmt.Sprint("`article_id` IN ?"), articleIds))
	buffer := dbr.NewBuffer()
	err := builder.Build(i.dialect, buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var tmp []*PArticle
	_, err = builder.LoadContext(ctx, &tmp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := make(map[string]string, len(articleIds))
	for _, p := range tmp {
		category, ok := categoryMap[p.Category]
		if ok {
			result[p.ArticleId] = category
		}
	}
	return result, nil
}

func (i *impl) GetNewUsers(ctx context.Context, date string) ([]string, error) {
	session := i.dbr.NewSession(nil)
	builder := session.Select("user_id").From("p_user").Where(dbr.Expr(fmt.Sprintf("`time_created` between ? and ?"), date+" 00:00:00", date+" 23:59:59"))
	buffer := dbr.NewBuffer()
	err := builder.Build(i.dialect, buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var result []string
	_, err = builder.LoadContext(ctx, &result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}
