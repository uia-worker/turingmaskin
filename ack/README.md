# Notater om Ackermann funksjonen
- Den vanskeligste funksjon å beregne?

Observasjon: m og n blir redusert med hver iterasjon, så det er mulig å si at hvis de blir redusert, så skal de bli 0 og funksjonen skal avslutte, selv om eksperimentelt kan det ikke påvises.

Med i,j lik 5 trenges det en udefinert (enorm) tid for utførelse av funksjonen.

ack(4,1) er 65533 og tar noen minutter for å beregne (noen år tilbake). Tiden for neste tall (4,2) kan estimeres til ca. 65 000 * 3 minutter dvs. rundt 4 måneder. Hvis vi øker i,j vesentlig, så kommer vi veldig fort opp i enormt lang tid. Problemet er beregnbart, men vi kan ikke bevise det for alle naturlige tall pga. den enorme tiden det ville tatt (mye mer enn sekunder siden Big Bang, som kan estimeres med 2 opphøyd i 500-600).

Ackermann funksjonen kaller opp seg selv, og Ackermann funksjonen, som har kalt opp en Ackermann funksjon, kaller opp seg selv igjen, slik at det blir en rekke av n-ordens funksjoner som kaller opp n-ordens funksjoner. En mulig assosiasjon er 2 opphøyd i 2 opphøyd i 2 .... uendelig mange ganger.
