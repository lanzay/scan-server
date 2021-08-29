import React from "react";

export default function Barcodes(props) {

    if (!props.barcodes) {
        return <>N/A</>
    }

    let list = props.barcodes.map((barcode, i) => {
        return (<div className={"row border-bottom font-monospace"} key={barcode.barcode}>
            <div className={"col-1"}>{i+1}</div>
            <div className={"col-10"}>{barcode.barcode}</div>
            <div className={"col-1"}>{barcode.count}</div>
        </div>)
    })

    return <div>
        <div className={"row border-bottom border-top bg-light"}>
            <div className={"col-1"}>№</div>
            <div className={"col-10"}>ШК</div>
            <div className={"col-1"}>Кол</div>
        </div>
        {list}
    </div>
}
