let ws;

$(() => {
    ws = new WebSocket('ws://localhost:80/clicker');
    ws.addEventListener('message', function(e) {
        const msg = JSON.parse(e.data);
        $('#count').text(msg.count)
    });
});

$('#counter').click(() => {
    ws.send(
        JSON.stringify({
            count: 1
        })
    );
})
