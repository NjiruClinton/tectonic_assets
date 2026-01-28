let cpuChartInstance = null;
let memoryChartInstance = null;
const processColors = [
    'rgb(59, 130, 246)',
    'rgb(16, 185, 129)',
    'rgb(245, 101, 101)',
    'rgb(251, 191, 36)',
    'rgb(139, 92, 246)',
    'rgb(236, 72, 153)',
    'rgb(14, 165, 233)',
    'rgb(34, 197, 94)',
    'rgb(249, 115, 22)',
    'rgb(168, 85, 247)'
]

document.addEventListener('DOMContentLoaded', () => {
    const now = new Date()
    const fiveHoursAgo = new Date(now - 5 * 60 * 60 * 1000)
    document.getElementById('cpuStartDate').value = formatDateForInput(fiveHoursAgo)
    document.getElementById('cpuEndDate').value = formatDateForInput(now)
    document.getElementById('memoryStartDate').value = formatDateForInput(fiveHoursAgo)
    document.getElementById('memoryEndDate').value = formatDateForInput(now)
})

function formatDateForInput(date) {
    return date.toISOString().slice(0, 16)
}
document.addEventListener('htmx:beforeSwap', (event) => {
    const targetId = event.detail.target.id

    if (targetId === 'cpuChartContainer' || targetId === 'memoryChartContainer') {
        event.detail.shouldSwap = false // Prevent the default swap - we'll handle it manually
        try {
            const data = JSON.parse(event.detail.xhr.responseText)
            if (targetId === 'cpuChartContainer') {
                event.detail.target.innerHTML = '<canvas id="cpuChart" width="800" height="400" class="w-full"></canvas>';
                setTimeout(() => {
                    renderCpuChart(data)
                    updateProcessStats('cpu', data)
                }, 100)
            } else if (targetId === 'memoryChartContainer') {
                event.detail.target.innerHTML = '<canvas id="memoryChart" width="800" height="400" class="w-full"></canvas>';
                setTimeout(() => {
                    renderMemoryChart(data)
                    updateProcessStats('memory', data)
                }, 100)
            }
        } catch (error) {
            console.error('Error parsing JSON data:', error)
            event.detail.target.innerHTML = showErrorHTML('Error loading data')
        }
    }
})
document.addEventListener('htmx:afterSwap', (event) => {
    const targetId = event.detail.target.id
    if (targetId === 'cpuChartContainer') {
        const data = JSON.parse(event.detail.xhr.responseText)
        renderCpuChart(data)
        updateProcessStats('cpu', data)
    } else if (targetId === 'memoryChartContainer') {
        const data = JSON.parse(event.detail.xhr.responseText)
        renderMemoryChart(data)
        updateProcessStats('memory', data)
  }
})
function renderCpuChart(data) {
    const ctx = document.getElementById('cpuChart')
    if (!ctx) return
    const processData = groupDataByProcess(data)

    if (cpuChartInstance) {
        cpuChartInstance.destroy()
  }

    const datasets = Object.keys(processData).map((processName, index) => ({
        label: processName,
        data: processData[processName].map(entry => ({
            x: new Date(entry.timestamp),
            y: entry.usage
        })),
        borderColor: processColors[index % processColors.length],
        backgroundColor: processColors[index % processColors.length] + '20',
        borderWidth: 2,
        fill: false,
        tension: 0.1
    }))

    cpuChartInstance = new Chart(ctx, {
        type: 'line',
        data: { datasets },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
                title: {
                    display: true,
                    text: 'CPU Usage Over Time'
                },
                legend: {
                    position: 'bottom',
                    labels: {
                        boxWidth: 12
                    }
                }
            },
            scales: {
                x: {
                    type: 'time',
                    time: {
                        unit: 'minute',
                        displayFormats: {
                            minute: 'HH:mm',
                            hour: 'HH:mm'
                        }
                    },
                    title: {
                        display: true,
                        text: 'Time'
                    }
                },
                y: {
                    beginAtZero: true,
                    title: {
                        display: true,
                        text: 'CPU Usage (%)'
                    },
                    ticks: {
                        callback: function(value) {
                            return value + '%';
        }
      }
                }
            },
            interaction: {
                intersect: false,
                mode: 'index'
    }
        }
    })
}

function renderMemoryChart(data) {
    const ctx = document.getElementById('memoryChart')
    if (!ctx) return
    const processData = groupDataByProcess(data)

    if (memoryChartInstance) {
        memoryChartInstance.destroy()
    }

    const datasets = Object.keys(processData).map((processName, index) => ({
        label: processName,
        data: processData[processName].map(entry => ({
            x: new Date(entry.timestamp),
            y: entry.usage / 1048576 // Convert bytes to MB
        })),
        borderColor: processColors[index % processColors.length],
        backgroundColor: processColors[index % processColors.length] + '20',
        borderWidth: 2,
        fill: false,
        tension: 0.1
    }))

    memoryChartInstance = new Chart(ctx, {
        type: 'line',
        data: { datasets },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
                title: {
                    display: true,
                    text: 'Memory Usage Over Time'
                },
                legend: {
                    position: 'bottom',
                    labels: {
                        boxWidth: 12
                    }
                }
            },
            scales: {
                x: {
                    type: 'time',
                    time: {
                        unit: 'minute',
                        displayFormats: {
                            minute: 'HH:mm',
                            hour: 'HH:mm'
                        }
                    },
                    title: {
                        display: true,
                        text: 'Time'
                    }
                },
                y: {
                    beginAtZero: true,
                    title: {
                        display: true,
                        text: 'Memory Usage (MB)'
                    },
                    ticks: {
                        callback: function(value) {
                            return value + ' MB';
                        }
                    }
                }
            },
            interaction: {
                intersect: false,
                mode: 'index'
            }
        }
    })
}

function groupDataByProcess(data) {
    const grouped = {}
    data.forEach(entry => {
        const processName = entry.process_name
        if (!grouped[processName]) {
            grouped[processName] = []
        }
        grouped[processName].push(entry)
    })
    Object.keys(grouped).forEach(processName => {
        grouped[processName].sort((a, b) => new Date(a.timestamp) - new Date(b.timestamp))
    })

    return grouped
}
function updateProcessStats(type, data) {
    const targetElement = type === 'cpu' ? 'topCpuProcesses' : 'topMemoryProcesses'
    const processStats = calculateProcessStats(data, type)
    const html = processStats.map((stat, index) => `
        <div class="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-900 rounded-md">
            <div class="flex items-center space-x-3">
                <div class="w-3 h-3 rounded-full" style="background-color: ${processColors[index % processColors.length]}"></div>
                <span class="font-medium text-gray-900 dark:text-gray-100">${stat.processName}</span>
            </div>
            <div class="text-right">
                <div class="text-sm font-semibold text-gray-900 dark:text-gray-100">${stat.avgUsage.toFixed(2)}${type === 'cpu' ? '%' : ' MB'}</div>
                <div class="text-xs text-gray-500 dark:text-gray-400">Max: ${stat.maxUsage.toFixed(2)}${type === 'cpu' ? '%' : ' MB'}</div>
            </div>
        </div>
    `).join('');

    const element = document.getElementById(targetElement);
    if (element) {
        element.innerHTML = html || '<div class="text-gray-500 dark:text-gray-400 text-center py-4">No data available</div>';
    }
}

function calculateProcessStats(data, type) {
    const processStats = {}

    data.forEach(entry => {
        const processName = entry.process_name
        if (!processStats[processName]) {
            processStats[processName] = {
                processName,
                totalUsage: 0,
                maxUsage: 0,
                count: 0
            }
        }
        const usage = type === 'memory' ? entry.usage / 1048576 : entry.usage
        processStats[processName].totalUsage += usage;
        processStats[processName].maxUsage = Math.max(processStats[processName].maxUsage, usage);
        processStats[processName].count++
    })

    return Object.values(processStats)
        .map(stat => ({
            ...stat,
            avgUsage: stat.totalUsage / stat.count
        }))
        .sort((a, b) => b.avgUsage - a.avgUsage)
        .slice(0, 5)
}

function showErrorHTML(message) {
    return `
        <div class="flex items-center justify-center p-8 text-red-600">
            <svg class="w-6 h-6 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            ${message}
        </div>
    `
}
