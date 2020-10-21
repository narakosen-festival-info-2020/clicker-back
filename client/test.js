let ws;

$(() => {
    ws = new WebSocket('ws://clicker-back-lb-465582205.ap-northeast-1.elb.amazonaws.com/clicker');
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
