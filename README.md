# Battery Management Services

## Circuit Diagram
Beräkna den utvecklade effekten i batteriets inre resistans Ri
Figur 1. U0=1.5V, I=0.2A, Ri=0.3Ohm

![image1]()

- Ui​=I⋅Ri
- The circuit is concatenation so that 'I' is consistent.
- P=I2⋅Ri​
- P=I2⋅Ri​=(0.2A)2⋅(0.3Ω)=0.04W

## Case study

![image2](https://ev-database.org/img/auto/BYD_DOLPHIN/BYD_DOLPHIN-01@2x.jpg)
- Fastcharge Power (10-80%): 65kW DC
- V2H via DC Supported: No
- Location: SE4 (nordpoolgroup.se)
- Residence: apartment 6.5kw 
- Role: a Student (from 8 am to 4 pm at school)

### Results: 

1. Scenario 1

| Max Power for the User (kW) | Energy Price Zone | Max PV Production | Daily PV Production | BESS Capacity | BESS CH | BESS DIS | EV Battery Capacity (kWh) | EV Battery SoC Morning | EV Battery SoC Evening | Daily Energy from Grid | Daily Energy Cost |
|-----------------------------|-------------------|-------------------|---------------------|---------------|---------|----------|--------------------------|------------------------|------------------------|------------------------|-------------------|
| 6.5                         | SE4               | n/a               | n/a                 | n/a           | n/a     | n/a      | 65                       | 56.38                  | 30.77                  | 73.27                  | 68.71             |

2. Scenario 2

| Max Power for the User (kW) | Energy Price Zone | Max PV Production | Daily PV Production | BESS Capacity | BESS CH & DIS | EV Battery Capacity | EV Battery SoC Morning | EV Battery SoC Evening | Daily Energy from Grid | Daily Energy Cost |
|------------------------------|-------------------|-------------------|---------------------|---------------|---------------|----------------------|------------------------|------------------------|------------------------|-------------------|
| 6.5                          | SE4               | 9.13              | 36.04               | n/a           | n/a           | 65                   | 78.37                  | 30.77                  | 56.78                  | 49.38             |

3. Scenario 3

| Max Power for the User (kW) | Energy Price Zone | Max PV Production | Daily PV Production | BESS Capacity (kWh) | BESS CH | BESS DIS | EV Battery Capacity (kWh) | EV Battery SoC Morning | EV Battery SoC Evening | Daily Energy from Grid | Daily Energy Cost |
|------------------------------|-------------------|-------------------|---------------------|----------------------|---------|----------|----------------------------|------------------------|------------------------|------------------------|-------------------|
| 6.5                          | SE4               | 9.13              | 36.04               | 18.3                 | 17.02   | 14.56    | 65                         | 75.67                  | 30.77                  | 52.22                  | 44.46             |