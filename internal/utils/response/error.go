package response

const (
	ErrBadRequest   = "bad_request"
	ErrUnauthorized = "unauthorized"
	ErrNotFound     = "not_found"
	ErrPermission   = "no_permission"
)

const (
	errBadRequest   = "yêu cầu không hợp lệ"
	errUnauthorized = "Không xác thực được người dùng"
	errNotFound     = "không tìm thấy"
	errPermission   = "không có quyền thực hiện hành động này"
)

const ()

type ErrorResponse struct {
	Key     string
	Message string
}

var (
	list []ErrorResponse
)

func Init() {
	list = make([]ErrorResponse, 0)

	defaultErrs := []ErrorResponse{
		{
			Key:     ErrBadRequest,
			Message: errBadRequest,
		},
		{
			Key:     ErrUnauthorized,
			Message: errUnauthorized,
		},
		{
			Key:     ErrNotFound,
			Message: errNotFound,
		},
		{
			Key:     ErrPermission,
			Message: errPermission,
		},
	}

	list = append(list, defaultErrs...)
}

func AddErrors(errs []ErrorResponse) {
	list = append(list, errs...)
}
