package school

import (
	"net/http"

	"github.com/hashmup/QuestionBankAPI/src/domain/service"
	"github.com/hashmup/QuestionBankAPI/src/interfaces"
)

type Dependency struct {
	SchoolService service.SchoolService
}

func (d *Dependency) GetSchoolsHandler(w http.ResponseWriter, r *http.Request) {
	schools, err := d.SchoolService.GetSchools(r.Context())
	if err != nil {
		res := interfaces.NewErrorResponse(http.StatusInternalServerError, err.Error())
		interfaces.Redererer.JSON(w, res.Status, res)
		return
	}

	res := encodeGetSchoolsResponse(schools)

	interfaces.Redererer.JSON(w, http.StatusOK, res)
}
