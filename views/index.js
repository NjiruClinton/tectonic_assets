let chartInstance = null;

document.addEventListener('htmx:beforeRequest', () => {
  document.getElementById('loader').classList.remove('hidden');
});

document.addEventListener('htmx:afterSwap', (event) => {
  if (event.detail.target.id === 'cpuUsageChart') {
    const data = JSON.parse(event.detail.xhr.responseText);
    renderChart(data);
    document.getElementById('loader').classList.add('hidden');
  }
});

function renderChart(data) {
  const ctx = document.getElementById('cpuUsageChart').getContext('2d');
  const labels = data.map(entry => entry.timestamp);
  const usageData = data.map(entry => entry.usage);

  if (chartInstance) {
    chartInstance.destroy();
  }

  chartInstance = new Chart(ctx, {
    type: 'line',
    data: {
      labels: labels,
      datasets: [{
        label: 'CPU Usage',
        data: usageData,
        borderColor: 'rgba(75, 192, 192, 1)',
        borderWidth: 1,
        fill: false
      }]
    },
    options: {
      scales: {
        x: {
          type: 'time',
          time: {
            unit: 'minute'
          }
        },
        y: {
          beginAtZero: true
        }
      }
    }
  });
}