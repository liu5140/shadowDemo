package service

import (
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	modelc "shadowDemo/shadow-framework/model"
)

type OrderService struct{}

var orderService *OrderService

func NewOrderService() *OrderService {
	if orderService == nil {
		l.Lock()
		if orderService == nil {
			orderService = &OrderService{}
		}
		l.Unlock()
	}
	return orderService
}

//创建
func (service *OrderService) CreateOrder(order *do.Order) (err error) {
	return model.GetModel().OrderDao.Create(order)
}

//通过id获取详情
func (service *OrderService) GetOrderByID(id int64) (order *do.Order, err error) {
	order.ID = id
	err = model.GetModel().OrderDao.Get(order)
	if err != nil {
		Log.Error(err)
		return order, err
	}
	return order, err
}

//通过id删除
func (service *OrderService) DeleteOrderByID(id int64) (err error) {
	if model.GetModel().OrderDao.Delete(&do.Order{ID: id}); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//通过id更新
func (service *OrderService) UpdateOrder(id int64, attrs map[string]interface{}) (err error) {
	if err = model.GetModel().OrderDao.Updates(id, attrs); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//查询
func (service *OrderService) SearchOrderPaging(condition *dao.OrderSearchCondition, pageNum int, pageSize int) (request []do.Order, count int, err error) {
	rowbound := modelc.NewRowBound(pageNum, pageSize)
	return service.searchOrder(condition, &rowbound)
}

func (service *OrderService) SearchOrderWithOutPaging(condition *dao.OrderSearchCondition) (request []do.Order, count int, err error) {
	return service.searchOrder(condition, nil)
}

func (service *OrderService) searchOrder(condition *dao.OrderSearchCondition, rowbound *modelc.RowBound) (request []do.Order, count int, err error) {
	result, count, err := model.GetModel().OrderDao.SearchOrders(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return result, count, err
}
