# Personal Assistant

# Project Beschrijving
Dit project is een digitale persoonlijke assistent die het dagelijks leven van mensen makkelijker kan maken. Enkele features van de assistent zijn: de gebruiker wakker maken, todo’s instellen, meldingen van afspraken geven en meer. Ook is het systeem gekoppeld aan meerdere slimme apparaten in huis. 
Het systeem zal voor het belangrijkste deel autonoom draaien en de gebruiker wakker maken wanneer dat echt nodig is. Dit wordt gerealiseerd door een connectie met verschillende externe bronnen. 
Het systeem wordt ontwikkeld in Golang. Dit is een programmeertaal ontwikkeld door Google. Een van de grote voordelen van Golang is de compatibiliteit, Golang is namelijk uitvoerbaar op een groot scala aan besturingssystemen. Verder is Golang heel licht in gebruik waardoor het draaien van het systeem mogelijk is op bijvoorbeeld een Raspberry Pi. 

De 1e uitwerking van dit project zal de Slimme Wekker zijn, aan de hand van deze basis kan het project verder uitgebreid worden tot een persoonlijke assistent die de gebruiker zijn leven op veel vlakken kan vereenvoudigen.

# Opbouw Project
Omdat het project in Golang gebouwd is verschilt de structuur met een applicatie in bijvoorbeeld Java of C#. Er wordt met packages gewerkt waarbij iedere package zijn eigen functionaliteit heeft waardoor deze dus deels afhankelijk van elkaar zijn. De basis van het systeem bevindt zich in het personalassistant package (van hieruit wordt de applicatie ook gestart, main.go). Verder is er de observer package gebouwt welke de communicatie tussen de verschillende triggers, events en actions afhandelt (meer hierover bij observer hoofdstuk). Tot slot zijn er de calendarsource trafficsource, alarmclock, philipshue en TestAlarmClock welke triggers of actions zijn.

# Observer
In dit pacakge wordt de communicatie tussen de verschillende triggers, events en actoins afgehandelt. Aan een event kunnen meerdere triggers gekoppeld worden. Wanneer zo'n trigger toegevoegd wordt aan een event wordt het event als listener bij de trigger aangemeld zodat het event aangeroepen kan worden wanneer de trigger "getriggerd" wordt. Op het event wordt dan een check gedaan of alle triggers zich al gemeldt hebben, als dat zo is dan roept het event de actions aan. Zij gaan dan bijvoorbeeld de lampen aanzetten of een muziekje afspelen.
