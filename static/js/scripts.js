document.getElementById('btnGetFinal').addEventListener('click', getFinalFunc)


async function getFinalFunc() {
    clearResponseBox()
    const averageMark = parseFloat(document.getElementById('averageMark').value.trim())
    const examMark = parseFloat(document.getElementById('examMark').value.trim())
    if (!averageMark && averageMark !== 0 || !examMark && examMark !== 0) {
        alert("Не все данные были введены")
        return
    }
    if (averageMark < 0.0 || averageMark > 100.0) {
        alert(`Средний балл ${averageMark} не является валидным баллом`)
        return
    }
    if (examMark < 0.0 || examMark > 100.0) {
        alert(`Экаминационный балл ${examMark} не является валидным баллом`)
        return
    }
    const data = {
        averageMark,
        examMark
    }
    const calculatorUrl = `${window.location.href}calculate`
    await fetch(calculatorUrl, 
        {
            method: 'POST',
            headers: {
        		'Content-Type': 'application/json'
			}, 
			body: JSON.stringify(data)
        })
        .then(response => {
            if (response.status === 400) {
                response.json().then(data => displayError(data))
                throw new Error("Do not proceed!")
            }
            response.json().then(data => processData(data))
        })
        .catch(err => console.log(`Error occured: ${err}`))
}

function displayError(data) {
    const responseBox = document.getElementById('response-box')
    responseBox.classList.add("bg-danger")
    responseBox.innerText = data.ErrorMessage
}

function processData(data) {
    const responseBox = document.getElementById('response-box')
    const finalMark = parseFloat(data.finalMark).toFixed(2)
    if (finalMark < 0 || finalMark > 100) {
        responseBox.classList.add("bg-danger")
        responseBox.innerText = "Были получены не валидные данные, проверьте корректность ввода"
        return
    }
    if (finalMark >= 70.0 && finalMark <= 100.0) {
        responseBox.classList.add("bg-success")
    } else if (finalMark >= 50.0 && finalMark < 70.0) {
        responseBox.classList.add("bg-warning")
    }else {
        responseBox.classList.add("bg-danger")
    }
    responseBox.innerText = finalMark
}

function clearResponseBox() {
    const responseBox = document.getElementById('response-box')
    responseBox.innerText = ""
    responseBox.classList.remove('bg-success')
    responseBox.classList.remove('bg-warning')
    responseBox.classList.remove('bg-danger')
}