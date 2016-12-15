let open = [], close = [];
let fun_1 = (map, start, stop) => {
    let cur;
    if (close.length === 0) cur = start;
    else cur = close[close.length - 1]
    if (cur.y - 1 >= 0 && map[cur.x][cur.y - 1]) {
        isHas(cur) && open.push({
            x: cur.x,
            y: cur.y - 1,
            H: Math.abs(stop.y - cur.y + 1) + Math.abs(stop.x - cur.x),
            G: Math.abs(cur.y - start.y) + Math.abs(cur.x - start.y),
        })
    }
    if (cur.y + 1 < map.length && map[cur.x][cur.y + 1]) {
        isHas(cur) && open.push({
            x: cur.x,
            y: cur.y + 1,
            H: Math.abs(stop.y - cur.y - 1) + Math.abs(stop.x - cur.x),
            G: Math.abs(cur.y - start.y) + Math.abs(cur.x - start.y),
        })
    }
    if (cur.x - 1 >= 0 && map[cur.x - 1][cur.y]) {
        isHas(cur) && open.push({
            x: cur.x - 1,
            y: cur.y,
            H: Math.abs(stop.y - cur.y) + Math.abs(stop.x - cur.x + 1),
            G: Math.abs(cur.y - start.y) + Math.abs(cur.x - start.y),
        })
    }
    if (cur.x + 1 < map[0].length && map[cur.x + 1][cur.y]) {
        isHas(cur) && open.push({
            x: cur.x + 1,
            y: cur.y,
            H: Math.abs(stop.y - cur.y) + Math.abs(stop.x - cur.x - 1),
            G: Math.abs(cur.y - start.y) + Math.abs(cur.x - start.y),
        })
    }
    cur = null;
    open.map(v => {
        if (!!cur) cur = v;
        else {
            if (cur.H + cur.G > v.H + v.G) {
                cur = v;
            }
        }
    })
    close.push(cur);
}

let isHas = (p) => {
    for (let i = 0; i <= close.length; i++) {
        if (close[i].x === p.x && close[i].y === p.y) {
            return false;
        }
    }
    return true;
}

let map = [[true, true], [false, true]], start = {x: 0, y: 0}, stop = {x: 1, y: 1};
fun_1(map, start, stop);