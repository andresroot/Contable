
     //--------------- autocompleta tercero
          var numeroPanel = 1;
        function autocompletaTercero(obj,elementoDv,elementoNombre) {

                $(document.body).on('focusout', obj, function (e) {

                    if ($(obj).val() == '') {
                    } else {

                        $(obj).val($(obj).val().replace('.', ''));

                        terceroCodigo = $(obj).val().replace('.', '');

                        var datosEnviar = {
                            "terceroCodigo": terceroCodigo
                        };
                        accion = "{% url 'terceroActual' %}";
                        $.ajax({
                            url: accion,
                            type: "POST",
                            async: false,
                            data: JSON.stringify(datosEnviar),
                            contentType: "application/json; charset=utf-8",
                            dataType: "json",
                            error: function (response) {
                                alert('No Existe Tercero');
                                $(obj).val('');
                            },
                            success: function (response) {
                                if (jQuery.isEmptyObject(response)) {

                                } else {
                                    $.each(response, function (i, item) {
                                        $(elementoDv).val(item.terceroDv);
                                        $(elementoNombre).val(item.terceroNombre);
                                    });
                                }
                            }
                        });
                    }
                });

                $(obj).autocomplete({
                    source: function (request, response) {
                        $.ajax({
                            url: "/terceroBuscar/",
                            type: 'GET',
                            dataType: "json",
                            data: {
                                search: request.term
                            },
                            success: function (data) {
                                response(data);
                            }
                        });
                    },
                    messages: {
                        noResults: '',
                        results: function () {
                        }
                    },
                    maxShowItems: 5,
                    response: function (event, ui) {

                        // if (ui.content.length === 0) {
                        $('#terceroNuevo').val($(obj).val());
                        ui.content.push({
                            label: " Crear Tercero :  " + $('#terceroNuevo').val(),
                            button: true
                        });
                        ///  } else {

                        ///    }
                    },
                    select: function (event, ui) {
                        var label = ui.item.label;
                        var value = ui.item.value;
                        valor = $(obj).val();
                        // alert(label);
                        //alert( $('#terceroNuevo').val());
                        valorBuscar = " Crear Tercero :  " + $('#terceroNuevo').val();

                        if (label == valorBuscar) {
                            valor = $('#terceroNuevo').val();
                            panelLista('terceroNuevo', 'True', valor, obj.replace('#', ''))

                        } else {
                            elemento = obj;

                            $('terceroNombre').html(ui.item.productoNombre);

                        }
                        //store in session
                    },
                    open: function (event, ui) {
                        // var d = $('.ui-autocomplete').append("<a href='/AdvancedSearch/[" + search_term + "]'>Crear Producto [" + search_term + "]</a>")
                    }
                });
            }
        window.document.addEventListener('myEvent', handleEvent, false)

        function handleEvent(e) {

            if (e.detail.valido == true) {
                //alert(e.detail.codigoElemento);

                if (e.detail.elementoPanel == "terceroCodigo") {
                    valor = e.detail.codigoElemento.replace('.', '');
                } else {
                    valor = e.detail.codigoElemento;
                }

                $('#' + e.detail.elementoPanel).val(valor);
                $('#' + e.detail.elementoPanel).focus();
                panelNuevo.close();
            } else {
                panelNuevo.close();
            }

            console.log(e.detail) // outputs: {foo: 'bar'}
        }

        function panelLista(modulo, panel, parametro, elemento) {
            numeroPanel = numeroPanel + 1;
            cadenaPanel = "panelph" + numeroPanel;
            url = '/' + modulo + '/' + panel + '/' + parametro + '/' + elemento;
            url = "<iframe src=\'" + url + "\' width=\'100%\' height=\'100%\' style=\'padding: 15px;\'></iframe>";
            panelNuevo = jsPanel.create({
                theme: {
                    bgContent: '#fff',
                    colorHeader: 'black',
                    border: '1px #A8A8A8 solid'
                },
                headerControls: {
                    maximize: 'remove',
                    size: 'xs'
                },
                id: cadenaPanel,
                size: {width: 800, height: 2200},
                contentSize: {width: '1400px', height: '780px'}, // must be object
                content: url,
                position: {
                    top: '350px',
                    left: '600px'
                },
                headerTitle: 'Sadconf Cloud 1.0'
            });
        }
        //---autocomppleta tercero
