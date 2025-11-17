package database

//import (
//	"context"
//	"gau-cloud-orchestrator/entity"
//
//	"gorm.io/gorm"
//)
//
//type VMRepository interface {
//	Create(ctx context.Context, vm *entity.VM) error
//	GetByID(ctx context.Context, id string) (*entity.VM, error)
//	GetByUserID(ctx context.Context, userID string) ([]*entity.VM, error)
//	Update(ctx context.Context, vm *entity.VM) error
//	Delete(ctx context.Context, id string) error
//	UpdateStatus(ctx context.Context, id string, status string) error
//}
//
//type vmRepository struct {
//	db *gorm.DB
//}
//
//func NewVMRepository(db *gorm.DB) VMRepository {
//	return &vmRepository{db: db}
//}
//
//func (r *vmRepository) Create(ctx context.Context, vm *entity.VM) error {
//	return r.db.WithContext(ctx).Create(vm).Error
//}
//
//func (r *vmRepository) GetByID(ctx context.Context, id string) (*entity.VM, error) {
//	var vm entity.VM
//	err := r.db.WithContext(ctx).Where("id = ?", id).First(&vm).Error
//	if err != nil {
//		return nil, err
//	}
//	return &vm, nil
//}
//
//func (r *vmRepository) GetByUserID(ctx context.Context, userID string) ([]*entity.VM, error) {
//	var vms []*entity.VM
//	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&vms).Error
//	if err != nil {
//		return nil, err
//	}
//	return vms, nil
//}
//
//func (r *vmRepository) Update(ctx context.Context, vm *entity.VM) error {
//	return r.db.WithContext(ctx).Save(vm).Error
//}
//
//func (r *vmRepository) Delete(ctx context.Context, id string) error {
//	return r.db.WithContext(ctx).Delete(&entity.VM{}, "id = ?", id).Error
//}
//
//func (r *vmRepository) UpdateStatus(ctx context.Context, id string, status string) error {
//	return r.db.WithContext(ctx).Model(&entity.VM{}).Where("id = ?", id).Update("status", status).Error
//}
