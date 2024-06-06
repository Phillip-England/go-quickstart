package component

import (
	"go-quickstart/internal/util"
)

func GuestHeader() string {
	return `
		<header class='p-4 flex flex-row border-b'>
			<div>
				<img src="/static/img/logo.svg" alt="logo" class="w-[150px]">
			</div>
		</header>
	`
}

func BaseForm(method string, action string, content string) string {
	formId := util.RandStr(12)
	return `
		<form id='` + formId + `' method="` + method + `" action="` + action + `" class='rounded flex flex-col gap-8'>
			` + content + `
		</form>
	`
}

func LoginForm(email string, password string, err string) string {
	return BaseForm("POST", "/", `
		`+FormName("Login")+`
		`+FormErr(err)+`
		`+TextInput("Email", "email", email)+`
		`+TextInput("Password", "password", password)+`
		`+FormSubmit("Login")+`
	`)
}

func TextInput(label string, name string, value string) string {
	inputId := util.RandStr(12)
	return `
		<div class='flex flex-col gap-2 text-sm'>
			<label for="` + name + `" class=''>` + label + `</label>
			<input id="` + inputId + `" value="` + value + `" type="text" name="` + name + `" id="` + name + `" class='border p-1 rounded focus:outline-none'>
		</div>
		<script>
			(() => {
				let label = qs('#` + inputId + `').previousElementSibling
				if (label.textContent === 'Password') {
					qs('#` + inputId + `').type = 'password'
					qs('#` + inputId + `').setAttribute('autocomplete', 'current-password')
					return
				}
				qs('#` + inputId + `').setAttribute('autocomplete', '` + label + `')
			})()
			qs('#` + inputId + `').addEventListener('input', function(e) {
				let form = Dom.climbUntil(e.target, (element) => {
					return element.tagName === 'FORM'
				})
				let err = form.querySelectorAll('.form-err')[0]
				err.classList.add('transition', 'duration-200', 'opacity-0')
			})
			qs('#` + inputId + `').addEventListener('focus', function(e) {
				e.target.classList.add('transition', 'duration-200', 'border-gray-500')
			})
			qs('#` + inputId + `').addEventListener('blur', function(e) {
				e.target.classList.remove('border-gray-500')
			})
		</script>
	`
}

func FormName(name string) string {
	return `
		<h2 class='text-2xl font-semibold'>` + name + `</h2>
	`
}

func FormErr(err string) string {
	return `
		<p class='form-err text-sm text-[#8B0000]'>` + err + `</p>
	`

}

func FormSubmit(label string) string {
	btnId := util.RandStr(12)
	labelId := util.RandStr(12)
	loaderId := util.RandStr(12)
	btnOverlayId := util.RandStr(12)
	return `
		<button id="` + btnId + `" type="submit" class='relative bg-primary text-white p-2 rounded text-sm flex items-center'>
			<p id="` + labelId + `" class='w-full'>` + label + `</p>
			<div id="` + btnOverlayId + `" class='absolute w-full h-full top-0 left-0 opacity-0 z-10 bg-black'></div>
			<div id="` + loaderId + `" class='absolute opacity-0 right-2 h-5 w-5 rounded-full z-20 border border-primary border-t-white animate-spin'></div>
		</button>
		<script>
			qs('#` + btnId + `').addEventListener('click', function() {
				Animus.fadeIn(qs('#` + loaderId + `'), 200)
				Animus.fadeIn(qs('#` + btnOverlayId + `'), 200, 0.5)
			});
		</script>
	`
}
