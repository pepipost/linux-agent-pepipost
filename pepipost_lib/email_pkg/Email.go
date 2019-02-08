/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package email_pkg

import "pepipost/pepipost_lib/models_pkg"
import "pepipost/pepipost_lib/configuration_pkg"

/*
 * Interface for the EMAIL_IMPL
 */
type EMAIL interface {
    CreateSendEmail (*string, *models_pkg.EmailBody) (*models_pkg.SendEmailResponse, error)
}

/*
 * Factory for the EMAIL interaface returning EMAIL_IMPL
 */
func NewEMAIL(config configuration_pkg.CONFIGURATION) *EMAIL_IMPL {
    client := new(EMAIL_IMPL)
    client.config = config
    return client
}