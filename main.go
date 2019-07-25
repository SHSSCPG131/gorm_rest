package ghorm_candidate_backened__code

import (
	app2 "ghorm_candidate_backened _code/app"
	config2 "ghorm_candidate_backened _code/config"
)

func main() {
	config:= config2.GetConfig()
	app := &app2.App{}
	app.Initialize(config)
	app.Run(":3000")
}
