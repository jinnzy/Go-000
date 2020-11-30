学习笔记 题目：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

回答：
  如老师上课所说，应该在我们的底层也就是`dao`层把`sql.ErrNoRows`这类错误转换为我们自己的错误，因为后面有可能换成`gorm` `redis` `mongo`等，让上层不强依赖`sql.ErrNoRows`，进行wrap返回上层。

伪代码

Dao层:
```Go
const (
  ErrNotFound = errors.New("not found")
)
type Dao struct {
  db     *sql.DB
}
func (d *dao) FindGoodsByID(id uint) (model.Goods, error) {
  var goods Goods
  err := d.db.QueryRow("select ... where id = ?", id).Scan(&goods)
  if errors.Is(err, sql.ErrNoRows) {
    err = ErrNotFound
  }
  return errors.Wrap(err, fmt.Sprintf("find goods by id: %v", err))
}
```
service层:
```Go

type Service struct {
  dao   *dao.Dao
}

// 在这层判断是否拿到值进行一些逻辑处理
func (s *Service) FindGoodsByID(id uint) (model.Goods, error) {
  goods,err := s.dao.FindGoodsByID(id)
  if err != nil && errors.Is(err, dao.ErrNotFound) {
    // 进行一些逻辑处理
    return ...
  }
  return ...
}







```
