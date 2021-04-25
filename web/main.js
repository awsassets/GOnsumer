$(document).ready(function() {
    var t = $('#ipTable').DataTable({
            "order": [
                [0, "desc"]
            ],
            "paging": false,
            "info": false,
            "searching": false,
            "columnDefs": [{
                "targets": -1,
                "data": null,
                "defaultContent": `<button type="button" class="btn btn-primary btn-sm">Check</button>`
            }]
        }),
        counter = 0,
        socket = new WebSocket("ws://127.0.0.1:1234/ws");

    socket.onmessage = function(event) {
        resolvedIP = resolveIP(event.data);
        if (resolvedIP !== null) {
            counter++;
            t.row.add([
                counter,
                resolvedIP.ip,
                resolvedIP.port,
                "Not checked",
                "ads"
            ]).draw(false);
        }
    };

    $('#ipTable tbody').on('click', 'button', function() {
        var row = t.row($(this).parents('tr')),
            data = row.data(),
            ip = data[1],
            port = data[2];
        $.get("/scan/" + ip + "/" + port, function(status) {
            if (status == "open") {
                data[3] = "Opened";
                row.invalidate().draw();

            } else {
                data[3] = "Closed";
                row.invalidate().draw();
            }
        });
        $(this).prop("disabled", true);
    });

    function resolveIP(data) {
        if (data.length >= 2 && data) {
            if (data.match(/^(?:(?:2[0-4]\d|25[0-5]|1\d{2}|[1-9]?\d)\.){3}(?:2[0-4]\d|25[0-5]|1\d{2}|[1-9]?\d)(?:\:(?:\d|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5]))?$/g)) {
                if (data.match(/:/g)) {
                    var split = data.split(':');
                    return {
                        ip: split[0],
                        port: split[1]
                    }
                }
                return null;
            }
        }
        return null;
    }
});