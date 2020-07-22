new Vue({
    el: '#app',

    data: {
        ws: null,
        score:'0',
        totalScore: '0'
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://'+ window.location.host +'/update-score');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            console.log(msg.kill);
            self.totalScore = msg.kill;
            var elem = document.getElementById('total-score');
        });
    },

    methods: {
        send: function() {
            console.log("total", this.totalScore)
            this.ws.send(
                JSON.stringify({
                    kill: $('<p>').html(this.totalScore).text()
                }
            ));
        },
    }
});