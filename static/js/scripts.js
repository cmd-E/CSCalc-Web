document.getElementById('btnGetFinal').addEventListener('click', getFinalFunc)
const responseBox = document.getElementById('response-box')
const calculatorUrl = "http://localhost:8080/calculate"

async function getFinalFunc() {
    const averageMark = document.getElementById('averageMark').value
    const examMark = document.getElementById('examMark').value
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
        }).then(response => response.json()).then(data => responseBox.innerText = data);
}