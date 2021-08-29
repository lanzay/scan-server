import React, {useEffect, useRef, useState} from "react";
import * as api from "../../api/api";
import Job from "./Job";
import Barcodes from "./Barcodes";

export default function Scan(props) {

    const [job, setJob] = useState({})
    const [barcode, setBarcode] = useState()
    const barcodeInput = useRef(null);

    useEffect(() => {
        barcodeInput.current.focus();
        api.getJob(props.match.params.job_id).then(response => {
            setJob(response.data)
        })
    }, [])

    function onScan(e) {
        api.scanBarcode(job.id, barcode).then(response => {
            console.log(response.data)
            setBarcode('')
            setJob(response.data)
            barcodeInput.current.focus()
        })
    }

    function handleKeyPress(target) {
        if (target.charCode === 13) {
            onScan()
        }
    }

    return (
        <div className={"row"}>
            <h2>SCAN</h2>
            <Job job={job}/>

            <label htmlFor="barcode" className="form-label">Штрихкод</label>
            <input
                type="text" id="barcode" className="form-control btn-lg" ref={barcodeInput}
                autoComplete={"off"} autoFocus={true}
                value={barcode} onInput={(v) => setBarcode(v.target.value)}
                onKeyPress={handleKeyPress}
            />
            <button onClick={onScan} type="submit" className="btn btn-primary btn-lg mt-3">Отправить</button>

            <div>
                <h5 className={"mt-5"}>Barcodes</h5>
                <Barcodes barcodes={job.barcodes}/>
            </div>


            <div>
                <h5 className={"mt-5"}>Barcodes detail</h5>
                <Barcodes barcodes={job.barcodes_detail}/>
            </div>
        </div>
    )
}
