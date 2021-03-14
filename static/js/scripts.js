document.getElementById('btnGetFinal').addEventListener('click', getFinalFunc)
const responseBox = document.getElementById('response-box')
const calculatorUrl = "http://localhost:8080/calculate"

async function getFinalFunc() {
    const averageMark = document.getElementById('averageMark').value
    const examMark = document.getElementById('examMark').value
    if (averageMark.trim() === "" || examMark.trim() === "") {
        alert("Не все данные были введены")
        return
    }
    const data = {
        averageMark,
        examMark
    }
    await fetch(calculatorUrl, 
        {
            method: 'POST',
            headers: {
        		'Content-Type': 'application/json'
			}, 
			body: JSON.stringify(data)
        }).then(response => response.json()).then(data => {
            const finalMark = parseFloat(data)
                if (finalMark >= 70.0 && finalMark <= 100.0) {
                    responseBox.classList.add("bg-success")
                } else if (finalMark >= 50.0 && finalMark <= 69.0) {
                    responseBox.classList.add("bg-warning")
                }else {
                    responseBox.classList.add("bg-danger")
                }
            responseBox.innerText = data
        });
}