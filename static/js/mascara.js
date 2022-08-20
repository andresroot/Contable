
function Coma(x) {
    return x.toString().replace(/\B(?<!\.\d*)(?=(\d{3})+(?!\d))/g, ",");
}

function formatomoneda(cnumero)
{
    //  12222222.12  a 12.222.222,12
    var parts = cnumero.toString().split(".");
    var num = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ".") + (parts[1] ? "," + parts[1] : "");
    if(num=="0,00")
    {
        num="";
    }
    return num;

}
function limpiarformato(cnumero) {
    // 12.222.222,12  a 12222222.12
    cnumero = cnumero.replace("$", "");
    cnumero=cnumero.replace(/\./g,'');
    cnumero = cnumero.replace(/\,/g,'.');
    return Number(cnumero)
}
function formatoguardar(cnumero) {
    // 12.222.222,12  a 12222222.12
    cnumero = cnumero.replace("$", "");
    cnumero=cnumero.replace(/\./g,'');
    cnumero = cnumero.replace(/\,/g,'.');
    return (cnumero)
}
