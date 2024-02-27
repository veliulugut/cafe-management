// Code generated by ent, DO NOT EDIT.

package menu

import (
	"cafe-management/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldID, id))
}

// MenuID applies equality check predicate on the "menu_id" field. It's identical to MenuIDEQ.
func MenuID(v int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldMenuID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldName, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldCategory, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldUpdatedAt, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldDescription, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldPrice, v))
}

// MenuImageURL applies equality check predicate on the "menu_image_url" field. It's identical to MenuImageURLEQ.
func MenuImageURL(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldMenuImageURL, v))
}

// MenuIDEQ applies the EQ predicate on the "menu_id" field.
func MenuIDEQ(v int) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldMenuID, v))
}

// MenuIDNEQ applies the NEQ predicate on the "menu_id" field.
func MenuIDNEQ(v int) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldMenuID, v))
}

// MenuIDIn applies the In predicate on the "menu_id" field.
func MenuIDIn(vs ...int) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldMenuID, vs...))
}

// MenuIDNotIn applies the NotIn predicate on the "menu_id" field.
func MenuIDNotIn(vs ...int) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldMenuID, vs...))
}

// MenuIDGT applies the GT predicate on the "menu_id" field.
func MenuIDGT(v int) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldMenuID, v))
}

// MenuIDGTE applies the GTE predicate on the "menu_id" field.
func MenuIDGTE(v int) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldMenuID, v))
}

// MenuIDLT applies the LT predicate on the "menu_id" field.
func MenuIDLT(v int) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldMenuID, v))
}

// MenuIDLTE applies the LTE predicate on the "menu_id" field.
func MenuIDLTE(v int) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldMenuID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldName, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldCategory, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldUpdatedAt, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldDescription, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldPrice, v))
}

// MenuImageURLEQ applies the EQ predicate on the "menu_image_url" field.
func MenuImageURLEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEQ(FieldMenuImageURL, v))
}

// MenuImageURLNEQ applies the NEQ predicate on the "menu_image_url" field.
func MenuImageURLNEQ(v string) predicate.Menu {
	return predicate.Menu(sql.FieldNEQ(FieldMenuImageURL, v))
}

// MenuImageURLIn applies the In predicate on the "menu_image_url" field.
func MenuImageURLIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldIn(FieldMenuImageURL, vs...))
}

// MenuImageURLNotIn applies the NotIn predicate on the "menu_image_url" field.
func MenuImageURLNotIn(vs ...string) predicate.Menu {
	return predicate.Menu(sql.FieldNotIn(FieldMenuImageURL, vs...))
}

// MenuImageURLGT applies the GT predicate on the "menu_image_url" field.
func MenuImageURLGT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGT(FieldMenuImageURL, v))
}

// MenuImageURLGTE applies the GTE predicate on the "menu_image_url" field.
func MenuImageURLGTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldGTE(FieldMenuImageURL, v))
}

// MenuImageURLLT applies the LT predicate on the "menu_image_url" field.
func MenuImageURLLT(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLT(FieldMenuImageURL, v))
}

// MenuImageURLLTE applies the LTE predicate on the "menu_image_url" field.
func MenuImageURLLTE(v string) predicate.Menu {
	return predicate.Menu(sql.FieldLTE(FieldMenuImageURL, v))
}

// MenuImageURLContains applies the Contains predicate on the "menu_image_url" field.
func MenuImageURLContains(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContains(FieldMenuImageURL, v))
}

// MenuImageURLHasPrefix applies the HasPrefix predicate on the "menu_image_url" field.
func MenuImageURLHasPrefix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasPrefix(FieldMenuImageURL, v))
}

// MenuImageURLHasSuffix applies the HasSuffix predicate on the "menu_image_url" field.
func MenuImageURLHasSuffix(v string) predicate.Menu {
	return predicate.Menu(sql.FieldHasSuffix(FieldMenuImageURL, v))
}

// MenuImageURLEqualFold applies the EqualFold predicate on the "menu_image_url" field.
func MenuImageURLEqualFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldEqualFold(FieldMenuImageURL, v))
}

// MenuImageURLContainsFold applies the ContainsFold predicate on the "menu_image_url" field.
func MenuImageURLContainsFold(v string) predicate.Menu {
	return predicate.Menu(sql.FieldContainsFold(FieldMenuImageURL, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Menu) predicate.Menu {
	return predicate.Menu(sql.NotPredicates(p))
}
