// This file is automatically generated by qtc from "formLayouts.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/layouts/formLayouts.qtpl:1
package layouts

//line views/templates/layouts/formLayouts.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//
// Обязательніе подключения для форм редактирования.

//line views/templates/layouts/formLayouts.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/layouts/formLayouts.qtpl:5
func StreamPutHead(qw422016 *qt422016.Writer) {
	//line views/templates/layouts/formLayouts.qtpl:5
	qw422016.N().S(`
    <link rel="stylesheet" href="/fields.css" type="text/css"  media="screen">
`)
//line views/templates/layouts/formLayouts.qtpl:7
}

//line views/templates/layouts/formLayouts.qtpl:7
func WritePutHead(qq422016 qtio422016.Writer) {
	//line views/templates/layouts/formLayouts.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/layouts/formLayouts.qtpl:7
	StreamPutHead(qw422016)
	//line views/templates/layouts/formLayouts.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line views/templates/layouts/formLayouts.qtpl:7
}

//line views/templates/layouts/formLayouts.qtpl:7
func PutHead() string {
	//line views/templates/layouts/formLayouts.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/layouts/formLayouts.qtpl:7
	WritePutHead(qb422016)
	//line views/templates/layouts/formLayouts.qtpl:7
	qs422016 := string(qb422016.B)
	//line views/templates/layouts/formLayouts.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/layouts/formLayouts.qtpl:7
	return qs422016
//line views/templates/layouts/formLayouts.qtpl:7
}

//line views/templates/layouts/formLayouts.qtpl:9
func StreamPutEnd(qw422016 *qt422016.Writer) {
	//line views/templates/layouts/formLayouts.qtpl:9
	qw422016.N().S(`
    <script src="/yii.js"></script>
    <script src="/yii.validation.js"></script>
    <script src="/yii.activeForm.js"></script>
    <script src="/bootstrap.js"></script>
    <script type="text/javascript">jQuery(document).ready(function () {
            jQuery('#initial_form').yiiActiveForm(
                [
                    {"id":"rooms-title",
                        "name":"title",
                        "container":".field-rooms-title",
                        "input":"#rooms-title",
                        "validate":
                            function (attribute, value, messages, deferred, $form) {
                    yii.validation.required(value, messages, {"message":"Уникальное название cannot be blank."});
                    yii.validation.string(value, messages, {"message":"Уникальное название must be a string.",
                        "max":255,"tooLong":"Уникальное название should contain at most 255 characters.","skipOnEmpty":1});}},
                    {"id":"rooms-offer_type","name":"offer_type","container":".field-rooms-offer_type",
                        "input":"#rooms-offer_type",
                        "validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages,
                            {"message":"Тип предложения must be a string.","skipOnEmpty":1});}},
                    {"id":"rooms-id_rooms_categories_list",
                        "name":"id_rooms_categories_list",
                        "container":".field-rooms-id_rooms_categories_list",
                        "input":"#rooms-id_rooms_categories_list",
                        "validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Категория номера must be an integer.","skipOnEmpty":1});}},{"id":"rooms-rooms_count","name":"rooms_count","container":".field-rooms-rooms_count","input":"#rooms-rooms_count","validate":function (attribute, value, messages, deferred, $form) {yii.validation.required(value, messages, {"message":"Количество номеров (этой категории) cannot be blank."});yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Количество номеров (этой категории) must be an integer.","skipOnEmpty":1});}},{"id":"rooms-main_points_count","name":"main_points_count","container":".field-rooms-main_points_count","input":"#rooms-main_points_count","validate":function (attribute, value, messages, deferred, $form) {yii.validation.required(value, messages, {"message":"Количество основных мест cannot be blank."});yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Количество основных мест must be an integer.","skipOnEmpty":1});}},{"id":"rooms-room_quota","name":"room_quota","container":".field-rooms-room_quota","input":"#rooms-room_quota","validate":function (attribute, value, messages, deferred, $form) {yii.validation.required(value, messages, {"message":"Квота номеров cannot be blank."});yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Квота номеров must be an integer.","skipOnEmpty":1});}}], []);
            jQuery('#first_step').yiiActiveForm([{"id":"roomsfirststep-square","name":"square","container":".field-roomsfirststep-square","input":"#roomsfirststep-square","validate":function (attribute, value, messages, deferred, $form) {yii.validation.required(value, messages, {"message":"Площадь cannot be blank."});yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Площадь must be an integer.","skipOnEmpty":1});}},{"id":"roomsfirststep-smoking_policy","name":"smoking_policy","container":".field-roomsfirststep-smoking_policy","input":"#roomsfirststep-smoking_policy","validate":function (attribute, value, messages, deferred, $form) {yii.validation.required(value, messages, {"message":"Курение в номере cannot be blank."});yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Курение в номере must be an integer.","skipOnEmpty":1});}},{"id":"roomsfirststep-id_accessibility_list","name":"id_accessibility_list","container":".field-roomsfirststep-id_accessibility_list","input":"#roomsfirststep-id_accessibility_list","validate":function (attribute, value, messages, deferred, $form) {yii.validation.required(value, messages, {"message":"Доступность cannot be blank."});yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Доступность must be an integer.","skipOnEmpty":1});}},{"id":"roomsfirststep-count_children_no_point","name":"count_children_no_point","container":".field-roomsfirststep-count_children_no_point","input":"#roomsfirststep-count_children_no_point","validate":function (attribute, value, messages, deferred, $form) {yii.validation.required(value, messages, {"message":"Количество детей без места cannot be blank."});yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Количество детей без места must be an integer.","skipOnEmpty":1});}},{"id":"roomsfirststep-count_additional_points","name":"count_additional_points","container":".field-roomsfirststep-count_additional_points","input":"#roomsfirststep-count_additional_points","validate":function (attribute, value, messages, deferred, $form) {yii.validation.required(value, messages, {"message":"Количество дополнительных мест cannot be blank."});yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Количество дополнительных мест must be an integer.","skipOnEmpty":1});}},{"id":"roomsfirststep-equipment_media","name":"equipment_media","container":".field-roomsfirststep-equipment_media","input":"#roomsfirststep-equipment_media","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Оснащение: Видео / аудио / интернет must be a string.","skipOnEmpty":1});}},{"id":"roomsfirststep-equipment_bathroom","name":"equipment_bathroom","container":".field-roomsfirststep-equipment_bathroom","input":"#roomsfirststep-equipment_bathroom","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Оснащение: Ванная комната must be a string.","skipOnEmpty":1});}},{"id":"roomsfirststep-equipment_electronics","name":"equipment_electronics","container":".field-roomsfirststep-equipment_electronics","input":"#roomsfirststep-equipment_electronics","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Оснащение: Электроника must be a string.","skipOnEmpty":1});}},{"id":"roomsfirststep-equipment_beds","name":"equipment_beds","container":".field-roomsfirststep-equipment_beds","input":"#roomsfirststep-equipment_beds","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Оснащение: Кровати must be a string.","skipOnEmpty":1});}},{"id":"roomsfirststep-current_bad_type_count","name":"current_bad_type_count","container":".field-roomsfirststep-current_bad_type_count","input":"#roomsfirststep-current_bad_type_count","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Количество кроватей данного типа в номере must be an integer.","skipOnEmpty":1});}},{"id":"roomsfirststep-equipment_furniture","name":"equipment_furniture","container":".field-roomsfirststep-equipment_furniture","input":"#roomsfirststep-equipment_furniture","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Оснащение: Мебель must be a string.","skipOnEmpty":1});}},{"id":"roomsfirststep-equipment_others","name":"equipment_others","container":".field-roomsfirststep-equipment_others","input":"#roomsfirststep-equipment_others","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Оснащение: Прочее must be a string.","skipOnEmpty":1});}},{"id":"roomsfirststep-outside_window_view","name":"outside_window_view","container":".field-roomsfirststep-outside_window_view","input":"#roomsfirststep-outside_window_view","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Внешняя территория и вид из окон must be a string.","skipOnEmpty":1});}},{"id":"roomsfirststep-outside_window_photos","name":"outside_window_photos","container":".field-roomsfirststep-outside_window_photos","input":"#roomsfirststep-outside_window_photos","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Внешняя территория и вид из окон: Фотографии must be a string.","skipOnEmpty":1});}},{"id":"roomsfirststep-equipment_design_specifics","name":"equipment_design_specifics","container":".field-roomsfirststep-equipment_design_specifics","input":"#roomsfirststep-equipment_design_specifics","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Оснащение: Специфика конструкции must be a string.","skipOnEmpty":1});}},{"id":"roomsfirststep-family_services","name":"family_services","container":".field-roomsfirststep-family_services","input":"#roomsfirststep-family_services","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Развлечения и семейные услуги must be a string.","skipOnEmpty":1});}}], []);
            jQuery('#second_step').yiiActiveForm([{"id":"roomssecondstep-base_price_per_night","name":"base_price_per_night","container":".field-roomssecondstep-base_price_per_night","input":"#roomssecondstep-base_price_per_night","validate":function (attribute, value, messages, deferred, $form) {yii.validation.required(value, messages, {"message":"Базовая цена за ночь cannot be blank."});yii.validation.number(value, messages, {"pattern":/^\s*[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?\s*$/,"message":"Базовая цена за ночь must be a number.","skipOnEmpty":1});}},{"id":"roomssecondstep-allowance_additional_point_adult","name":"allowance_additional_point_adult","container":".field-roomssecondstep-allowance_additional_point_adult","input":"#roomssecondstep-allowance_additional_point_adult","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?\s*$/,"message":"Размер доплаты за использование одного дополнительного места: Взрослым must be a number.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_for_less_main_point","name":"discount_for_less_main_point","container":".field-roomssecondstep-discount_for_less_main_point","input":"#roomssecondstep-discount_for_less_main_point","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Снижаете ли вы цену при размещении менее N гостей на основных местах? must be an integer.","skipOnEmpty":1});}},{"id":"roomssecondstep-minimal_price_use_for","name":"minimal_price_use_for","container":".field-roomssecondstep-minimal_price_use_for","input":"#roomssecondstep-minimal_price_use_for","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Минимальная цена распространяется на must be an integer.","skipOnEmpty":1});}},{"id":"roomssecondstep-set_value_yourself","name":"set_value_yourself","container":".field-roomssecondstep-set_value_yourself","input":"#roomssecondstep-set_value_yourself","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Указать значиния вручную must be an integer.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_format","name":"discount_format","container":".field-roomssecondstep-discount_format","input":"#roomssecondstep-discount_format","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Размер скидки (формат) must be a string.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_value","name":"discount_value","container":".field-roomssecondstep-discount_value","input":"#roomssecondstep-discount_value","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?\s*$/,"message":"Размер скидки (значение) must be a number.","skipOnEmpty":1});}},{"id":"roomssecondstep-price_with_discount","name":"price_with_discount","container":".field-roomssecondstep-price_with_discount","input":"#roomssecondstep-price_with_discount","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?\s*$/,"message":"Цена со скидкой (значения) must be a number.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_for_duration","name":"discount_for_duration","container":".field-roomssecondstep-discount_for_duration","input":"#roomssecondstep-discount_for_duration","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Предоставляете ли Вы скидки по продолжительности? must be an integer.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_duration_min_days","name":"discount_duration_min_days","container":".field-roomssecondstep-discount_duration_min_days","input":"#roomssecondstep-discount_duration_min_days","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Минимум дней must be an integer.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_duration_value","name":"discount_duration_value","container":".field-roomssecondstep-discount_duration_value","input":"#roomssecondstep-discount_duration_value","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?\s*$/,"message":"Размер скидки (значение) must be a number.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_duration_format","name":"discount_duration_format","container":".field-roomssecondstep-discount_duration_format","input":"#roomssecondstep-discount_duration_format","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Размер скидки (формат) must be a string.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_duration_use_for_day","name":"discount_duration_use_for_day","container":".field-roomssecondstep-discount_duration_use_for_day","input":"#roomssecondstep-discount_duration_use_for_day","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Применять для must be a string.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_duration_starting_day","name":"discount_duration_starting_day","container":".field-roomssecondstep-discount_duration_starting_day","input":"#roomssecondstep-discount_duration_starting_day","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Точка отсчета скидки must be an integer.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_duration_use_for_type","name":"discount_duration_use_for_type","container":".field-roomssecondstep-discount_duration_use_for_type","input":"#roomssecondstep-discount_duration_use_for_type","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Скидка действует для must be a string.","skipOnEmpty":1});}},{"id":"roomssecondstep-min_booking_time","name":"min_booking_time","container":".field-roomssecondstep-min_booking_time","input":"#roomssecondstep-min_booking_time","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"MIN срок бронирования must be an integer.","skipOnEmpty":1});}},{"id":"roomssecondstep-description","name":"description","container":".field-roomssecondstep-description","input":"#roomssecondstep-description","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Описание must be a string.","skipOnEmpty":1});}},{"id":"roomssecondstep-use_discount_on_children_main_place","name":"use_discount_on_children_main_place","container":".field-roomssecondstep-use_discount_on_children_main_place","input":"#roomssecondstep-use_discount_on_children_main_place","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[+-]?\d+\s*$/,"message":"Снижаете ли вы цену при размещении детей на основных местах? must be an integer.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_on_children_main_place_value","name":"discount_on_children_main_place_value","container":".field-roomssecondstep-discount_on_children_main_place_value","input":"#roomssecondstep-discount_on_children_main_place_value","validate":function (attribute, value, messages, deferred, $form) {yii.validation.number(value, messages, {"pattern":/^\s*[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?\s*$/,"message":"Размер скидки (значение) must be a number.","skipOnEmpty":1});}},{"id":"roomssecondstep-discount_on_children_main_place_format","name":"discount_on_children_main_place_format","container":".field-roomssecondstep-discount_on_children_main_place_format","input":"#roomssecondstep-discount_on_children_main_place_format","validate":function (attribute, value, messages, deferred, $form) {yii.validation.string(value, messages, {"message":"Размер скидки (формат) must be a string.","skipOnEmpty":1});}}], []);
        });
    </script>
`)
//line views/templates/layouts/formLayouts.qtpl:39
}

//line views/templates/layouts/formLayouts.qtpl:39
func WritePutEnd(qq422016 qtio422016.Writer) {
	//line views/templates/layouts/formLayouts.qtpl:39
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/layouts/formLayouts.qtpl:39
	StreamPutEnd(qw422016)
	//line views/templates/layouts/formLayouts.qtpl:39
	qt422016.ReleaseWriter(qw422016)
//line views/templates/layouts/formLayouts.qtpl:39
}

//line views/templates/layouts/formLayouts.qtpl:39
func PutEnd() string {
	//line views/templates/layouts/formLayouts.qtpl:39
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/layouts/formLayouts.qtpl:39
	WritePutEnd(qb422016)
	//line views/templates/layouts/formLayouts.qtpl:39
	qs422016 := string(qb422016.B)
	//line views/templates/layouts/formLayouts.qtpl:39
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/layouts/formLayouts.qtpl:39
	return qs422016
//line views/templates/layouts/formLayouts.qtpl:39
}
