document.getElementById('btnGetFinal').addEventListener('click', getFinalFunc)

async function getFinalFunc() {
    console.log('click')
    const calculatorUrl = "http://localhost:8080/calculate"
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
        }).then(response => response.json()).then(data => console.log(data));
}