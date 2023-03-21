go get -u github.com/go-chi/chi/v5
go get -u github.com/mattn/go-sqlite3

fetch('/apply', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    personal_id: '14',
    name: 'Super Mario',
    amount: '20000',
    term: '6',
  })
})
.then(response => {
  if (!response.ok) {
    throw new Error('Network response was not ok');
  }
  return response.json();
})
.then(data => {
  console.log(data);
})
.catch(error => {
  console.error('There was a problem with the fetch operation:', error);
});

curl -X POST localhost:8080/apply -d 'personal_id=14&name=Super+Mario&amount=500&term=6'