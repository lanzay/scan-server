import React, {useState} from "react";
import * as api from "../../api/api";

export default function Scan() {

    const [barcode, setBarcode] = useState()

    function onScan(e) {
        api.scan('qqq', barcode).then(response => {
            console.log(response.data)
            //setBarcode([response.data])
        })
    }

    return <div className={"row"}>
        <h2>SCAN</h2>
        <label htmlFor="barcode" className="form-label">Штрихкод</label>
        <input
            type="text" id="barcode" className="form-control btn-lg"
            value={barcode} onInput={(v) => setBarcode(v.target.value)}
        />
        <button onClick={onScan} type="submit" className="btn btn-primary btn-lg mt-3">Отправить</button>
    </div>
}
