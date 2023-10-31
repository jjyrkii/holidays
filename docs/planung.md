# Projektplanung Verwaltung von Personalressourcen (Kessel & Pfeiffer)

Die effiziente Verwaltung von Personalressourcen und Urlaubsanträgen ist für jedes Unternehmen von entscheidender Bedeutung. In diesem Projekt setzen wir das Ziel, die Abläufe in der Personal- und Urlaubsverwaltung, die bisher händisch erfolgten, zu optimieren. Durch die Einführung eines Systems zur Personal- und Urlaubsverwaltung werden Mitarbeiter und Vorgesetzte in der Urlaubsplanung und die Personalabteilung im Prozess der Personal- und Urlaubsverwaltung unterstützt.

## Veranschaulichung der Projektziele und Abläufe mittels UML-Diagrammen

Die folgenden Diagramme wurden entwickelt, um die Abläufe und Interaktionen im Rahmen unserer Projektplanung für die Personal- und Urlaubsverwaltung eines mittelständischen Unternehmens zu veranschaulichen. Diese visuellen Darstellungen dienen dazu, die verschiedenen Prozesse und Schnittstellen in unserem Projekt zu verdeutlichen. Wir haben Use-Case-Diagramme und Aktivitätsdiagramme verwendet, um die funktionalen Anforderungen und den Ablauf der Genehmigung und Weiterleitung von Urlaubsanträgen darzustellen.

### Anwendungsfälle

Anwendungsfall 1: Ein Mitarbeiter soll sich in das System einloggen können, um auf personalbezogene Funktionen zuzugreifen. Dazu gehört es Urlaubsanträge zu erstellen und einzureichen. Der eingereichte Urlaubsantrag wird an den Vorgesetzten zur Genehmigung gesendet. Die Vorgesetzten können eingegangene Urlaubsanträge anzeigen, genehmigen oder ablehnen. Die Mitarbeiter können ihren Urlaubsverbrauch und die verbleibenden Urlaubstage einsehen.

~~~plantuml
@startuml
left to right direction
actor Mitarbeiter as "Mitarbeiter"
actor Vorgesetzter as "Vorgesetzter"
Vorgesetzter --|> Mitarbeiter
rectangle "Personal- und Urlaubsverwaltung" {
    Mitarbeiter --> (Einloggen)
    Mitarbeiter --> (Urlaubsantrag erstellen)
    Mitarbeiter --> (Urlaubsantrag senden)
    Mitarbeiter --> (Urlaubsanträge anzeigen)
    Vorgesetzter --> (Urlaubsantrag genehmigen/ablehnen)
}
@enduml
~~~

Anwendungsfall 2: Die Personalabteilung kann Urlaubsanträge anzeigen und bearbeiten. Sie kann auch den Urlaubsstatus aktualisieren, z.B. um Urlaubstage zu subtrahieren oder wiederherzustellen. Außerdem kann die Personalabteilung Berichte über den Urlaubsverbrauch, verbleibende Urlaubstage und andere relevante Statistiken generieren.

~~~plantuml
left to right direction
actor Personalabteilung
rectangle "Personal- und Urlaubsverwaltung" {
    Personalabteilung --> (Urlaubsanträge anzeigen)
    Personalabteilung --> (Urlaubsstatus aktualisieren)
    Personalabteilung --> (Berichte generieren)
}
~~~

### Aktivitätdiagramm

Das folgende Aktivitätsdiagramm soll dabei helfen, den Genehmigungsprozess für Urlaubsanträge und die Weiterleitung an die Personalabteilung zu verstehen:

~~~plantuml
@startuml

start

:Mitarbeiter erstellt Urlaubsantrag;
:Mitarbeiter sendet Urlaubsantrag;

if (Vorgesetzter genehmigt Urlaubsantrag?) then (Ja)
  :Vorgesetzter genehmigt Urlaubsantrag;
  :Vorgesetzter leitet Urlaubsantrag an Personalabteilung weiter;
  :Personalabteilung aktualisiert Urlaubsstatus;
else (Nein)
  :Vorgesetzter lehnt Urlaubsantrag ab;
endif

:Urlaubsprozess abgeschlossen;
stop

@enduml
~~~


## TODO

+ Klassendiagramme
+ Technisch und organisatorische Maßnahmen dokumentieren
+ Architektur und Entwurfsmuster z.B. MVC
+ Objektrelationale Abbildung (ORM-Frameworks)
+ Relationale Datenbanken
+ Representational State Transfer (REST)
+ Serialisierung (Export/Import) von Daten mit JSON oder XML
+ Software-Dokumentationswerkzeuge
