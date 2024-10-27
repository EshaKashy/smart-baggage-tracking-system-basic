document.addEventListener('DOMContentLoaded', function () {
    fetch('http://localhost:8080/luggage')
        .then(response => response.json())
        .then(data => {
            const container = document.getElementById('luggage-container');
            data.forEach(luggage => {
                const luggageDiv = document.createElement('div');
                luggageDiv.className = 'luggage-item';
                luggageDiv.innerHTML = `
                    <h3>Luggage ID: ${luggage.id}</h3>
                    <p>Owner: ${luggage.owner}</p>
                    <p>Status: ${luggage.status}</p>
                    <p>Location: ${luggage.location}</p>
                `;
                container.appendChild(luggageDiv);
            });
        })
        .catch(error => console.error('Error fetching luggage data:', error));
});
