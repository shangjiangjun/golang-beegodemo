package database

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

/*
Int8, 等于Byte, 占1个字节.
Int16, 等于short, 占2个字节. -32768 32767
Int32, 等于int, 占4个字节. -2147483648 2147483647
Int64, 等于long, 占8个字节. -9223372036854775808 9223372036854775807

GO												MySQL

int, int32-设置auto或者名称为Id  					integer AUTO_INCREMENT
int64-设置auto或者名称为Id							bigint AUTO_INCREMENT
uint, uint32 - 设置 auto 或者名称为 Id				integer unsigned AUTO_INCREMENT
uint64 - 设置 auto 或者名称为 Id					bigint unsigned AUTO_INCREMENT
bool											bool
string - 默认为 size 255							varchar(size)
string - 设置 type(text)							longtext
time.Time-设置 type为date						date
time.Time										datetime
byte											tinyint unsigned
rune											integer
int												integer
int8											tinyint
int16											smallint
int32											integer
int64											bigint
uint											integer unsigned
uint8											tinyint unsigned
uint16											smallint unsigned
uint32											integer unsigned
uint64											bigint unsigned
float32											double precision
float64											double precision
float64 设置digits,decimals						numeric(digits, decimals)

orm
|pk|设置该字段为主键|
|auto|这只该字段自增，但是要求该字段必须为整型|
| default(0)| 设置该字段的默认值，需要注意字段类型和默认值类型一致|
|size(100) |设置该字段长度为100个字节，一般用来设置字符串类型 |
| null | 设置该字段允许为空，默认不允许为空|
| unique | 设置该字段全局唯一 |
| digits(12);decimals(4)| 设置浮点数位数和精度。比如这个是说，浮点数总共12位，小数位为四位。|
| auto_now| 针对时间类型字段，作用是保存数据的更新时间|
|auto_now_add| 针对时间类型字段,作用是保存数据的添加时间|
> 注意：当模型定义里没有主键时，符合int, int32, int64, uint, uint32, uint64 类型且名称为 Id 的 Field 将 被视为主键，能够自增. "


查询操作
qs.Filter("id", 1) // WHERE id = 1
qs.Filter("user__id",1) //where user.id =1
qs.Filter("id_in",10，20) //where age in(10,20)
qs.Filter("id__gte",18) //where id>=18
qs.Filter("id__gt",18) //where id>18
qs.Filter("id__,5) //where id<5
//gt 是greater缩写即大于
//get是 Greater or equal的缩写即大于等于
//lt 是less than 即小于
*/
type Users struct {
	Id        int64     `json:"id" orm:"column(id);pk;auto"`
	Name      string    `json:"name" orm:"column(name);size(15)"`
	Mobile    string    `json:"mobile" orm:"column(mobile);size(11)"`
	Avatar    string    `json:"avatar" orm:"column(avatar);size(191);null" description:"用户头像"`
	Password  string    `json:"password" orm:"column(password);size(191)"`
	Sign      bool      `json:"sign" orm:"column(sign);size(1);default(1)"`
	Status    bool      `json:"status" orm:"column(status);size(1);default(1)"`
	CreatedAt time.Time `json:"created_at" orm:"column(created_at);auto_now_add;type(datetime)"` //发布时间
	UpdatedAt time.Time `json:"updated_at" orm:"column(updated_at);auto_now;type(datetime)"`     //修改时间
}

// 返回表名
func (t *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

// 新增，成功返回主键ID
func AddUser(u *Users) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(u)
	return
}

// 根据主键ID查询数据
func GetUserById(id int64) (user *Users, err error) {
	o := orm.NewOrm()
	user = &Users{Id: id}
	if err = o.Read(user); err == nil {
		return user, nil
	}
	return nil, err
}

// 根据条件查询数据
func GetAllUsers(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (lists []interface{}, err error) {

	o := orm.NewOrm()
	// select 的query 语句
	qs := o.QueryTable(new(Users))
	// query k => v
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 对于每个排序字段，都有一个关联的顺序
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("排序方式必须是 'asc' 或 'desc'")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 只有一个顺序，所有已排序的字段都将按此顺序排序
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("排序方式必须是 ‘asc‘ 或 ‘desc‘")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("‘排序方式‘ 和 ‘排序字段’长度不匹配或 ‘order‘ 长度不是1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("未使用排序字段")
		}
	}

	var users []Users
	qs = qs.OrderBy(sortFields...).Offset(offset).Limit(limit)
	if _, err = qs.All(&users, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range users {
				lists = append(lists, v)
			}
		} else {
			for _, v := range users {
				list := make(map[string]interface{})
				// 反射： reflect.ValueOf() 函数获得值的反射值对象函数获得值的反射值对象
				item := reflect.ValueOf(v)
				for _, fname := range fields {
					list[fname] = item.FieldByName(fname).Interface()
				}
				lists = append(lists, list)
			}
		}
		return lists, nil
	}
	return nil, err
}

// 更新, 可更改为user的key
func UpdateUserById(user *Users) (err error) {
	o := orm.NewOrm()
	info := Users{Id: user.Id}

	if err = o.Read(&info); err == nil {
		var num int64
		if num, err = o.Update(user); err == nil {
			fmt.Println("数据已被修改，条数：", num)
		}
	}
	return
}

// 删除，
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	info := Users{Id: id}
	// 判断是否存在
	if err = o.Read(&info); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("已删除数据，条数：", num)
		}
	}
	return
}
