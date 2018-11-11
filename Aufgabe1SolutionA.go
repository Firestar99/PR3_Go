package main

/*
a.)
Dieser Fehler tritt auf, wegen des blockierenden Verhaltens der Schreib/Lesezugriffe auf den Channel.
Der Code läuft hier innerhalb eines Threads Zeile für Zeile nacheinander. Die Operation des Schreibens in den Kanal (c <- 42) blockiert die Ausführung des gesamten Programms, da Schreiboperationen auf einem synchronen Kanal nur erfolgreich sein können, wenn ein Empfänger bereit ist, diese Daten zu erhalten. Und wir erstellen den Empfänger nur in der nächsten Zeile.
*/
