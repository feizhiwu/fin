package fin

type Display struct {
	*Context
	Render
	Status int
	funcs  MF
	checks struct {
		verify  bool
		actions []string
	}
}

func (d *Display) Get(f func()) {
	d.funcs["GET-"+GetFuncName(f)] = f
}

func (d *Display) Post(f func()) {
	d.funcs["POST-"+GetFuncName(f)] = f
}

func (d *Display) Put(f func()) {
	d.funcs["PUT-"+GetFuncName(f)] = f
}

func (d *Display) Delete(f func()) {
	d.funcs["DELETE-"+GetFuncName(f)] = f
}

// checkAction 行为检查
func (d *Display) checkAction(value string) bool {
	d.Status = StatusWarn
	action := d.GetHeader("action")
	if action != "" && action == value {
		d.Status = StatusOK
		return true
	}
	return false
}

// Show 统一输出api数据
func (d *Display) Show(mix interface{}) {
	if d.Status == StatusInit {
		d.Status = StatusOK
	}
	//默认json格式
	if d.Render == nil {
		d.Render = Json(d.Context)
	}
	d.Render.Output(mix)
}

// Validate 参数检测
func (d *Display) Validate(val map[int]string, data map[string]interface{}) {
	d.Status = StatusOK
	for k, v := range val {
		if data[v] == nil {
			panic(k)
		}
	}
}

// HasKey 检测更新主键是否为空
func (d *Display) HasKey(data map[string]interface{}) {
	if data["id"] == nil {
		panic(80001)
	}
}

func (d *Display) CheckLogin(actions []string, verify bool) {
	d.checks.verify = verify
	d.checks.actions = actions
}

// ForceLogin 检测是否登录
func (d *Display) ForceLogin() {
	if d.Params["login_uid"] == nil {
		panic(80003)
	}
}

func (d *Display) Run() {
	action := d.GetHeader("action")
	f := d.funcs[d.Method+"-"+action]
	if f != nil {
		if d.checks.verify && InArray(len(d.checks.actions), func(i int) bool {
			return d.checks.actions[i] == action
		}) {
			d.ForceLogin()
			f()
		} else if !d.checks.verify && !InArray(len(d.checks.actions), func(i int) bool {
			return d.checks.actions[i] == action
		}) {
			d.ForceLogin()
			f()
		} else {
			f()
		}
	} else {
		d.Show(StatusWarn)
	}
}
