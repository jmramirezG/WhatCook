# Milestones

- [ ] [[M1] Modelo para almacenar los datos relevantes de los clientes.](https://github.com/jmramirezG/WhatCook/milestone/1)
  
>Según lo que hemos hablado con Jesús, él sólo necesita los nombres, pero si también quiere que se envíen los documentos automáticamente, necesitaremos los correos electrónicos.
>
>Tras conversaciones con Jesús, Ramona y Águeda, concluimos que será suficiente con guardar nombre completo y correo electrónico. Con esto se dará el hito por resuelto.

- [ ] [[M2] Método para generar documentos dada una lista de nombres y la plantilla del documento.](https://github.com/jmramirezG/WhatCook/milestone/2)

>Jesús necesita que cuando él ponga la lista de nombre y el tipo de documento que quiere (nos dará la plantilla) salgan los documentos con los datos rellenos.
>
>Este hito se considerará resuelto cuando los documentos salgan debidamente rellenos, con los nombres correctos y en archivos separados con un nombre del tipo \<NombrePersona>_\<TipoDocumento>.pdf.

- [ ] [[M3] Función de organización de documentos según persona a la que van dirigido y tipo de documento.](https://github.com/jmramirezG/WhatCook/milestone/3)

>Ramona no quiere organizar los documentos a mano. Hay que buscar una forma de que los documentos se organicen solos, asignándoles una etiqueta para que sea localizable posteriormente.
>
>Este hito se considerará resuelto cuando los documentos queden organizados en una base de datos, con etiquetas que indiquen la persona a la que iban dirigidos y el tipo de documento.

- [ ] [[M4] Modelo de visualización de documentos filtrados según criterios.](https://github.com/jmramirezG/WhatCook/milestone/4)

>Águeda no quiere depender de cómo Ramona ordene los documentos. Tenemos que proporcionar una forma en la que filtrar los documentos, donde Águeda pueda elegir filtros para visualizarlos como ella quiera.
>
>Este hito se considerará resuelto cuando tengamos una forma unificada de poder acceder a la base de datos de documentos del hito [M3] desde varios dispositivos simultáneamente. Se deberá poder filtrar por persona, fecha y tipo de documento (según los criterios que nos ha dado Águeda).

- [ ] [[M5] Función de envío automático de los documentos generados a sus destinatarios.](https://github.com/jmramirezG/WhatCook/milestone/5)

>Jesús quiere que cuando genere los documentos, se envíen a sus respectivos destinatarios.
>
>Este hito estará resuelto cuando, una vez generados los documentos (milestone [M2]), se puedan enviar a sus destinatarios de forma automática junto a un mensaje de correo personalizado de la forma: "Saludos \<NombrePersona>, le envío su \<TipoDocumento>".
