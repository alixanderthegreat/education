"""Functions to prevent a nuclear meltdown."""


def is_criticality_balanced(temperature, neutrons_emitted):
    """Verify criticality is balanced.

    :param temperature: int or float - temperature value in kelvin.
    :param neutrons_emitted: int or float - number of neutrons emitted per second.
    :return: bool - is criticality balanced?

    A reactor is said to be critical if it satisfies the following conditions:
    - The temperature is less than 800 K.
    - The number of neutrons emitted per second is greater than 500.
    - The product of temperature and neutrons emitted per second is less than 500000.
    """

    return temperature < 800 and neutrons_emitted > 500 and (temperature * neutrons_emitted) < 500000


def reactor_efficiency(voltage, current, theoretical_max_power):
    """Assess reactor efficiency zone.

    :param voltage: int or float - voltage value.
    :param current: int or float - current value.
    :param theoretical_max_power: int or float - power that corresponds to a 100% efficiency.
    :return: str - one of ('green', 'orange', 'red', or 'black').

    Efficiency can be grouped into 4 bands:

    1. green -> efficiency of 80% or more,
    2. orange -> efficiency of less than 80% but at least 60%,
    3. red -> efficiency below 60%, but still 30% or more,
    4. black ->  less than 30% efficient.

    The percentage value is calculated as
    (generated power/ theoretical max power)*100
    where generated power = voltage * current
    """
    generated = voltage * current
    value = (generated / theoretical_max_power) * 100
    efficiency = ""

    green = value >= 80
    low = value >= 60
    high = value < 80
    orange = low and high 
    low = value < 60
    high = value >= 30
    red = low and high 
    black = value < 30

    if green:
        efficiency = "green"
    elif orange:
        efficiency = "orange"
    elif red:
        efficiency = "red"
    elif black:
        efficiency = "black"

    return efficiency

def fail_safe(temperature, neutrons_produced_per_second, threshold):
    """Assess and return status code for the reactor.

    :param temperature: int or float - value of the temperature in kelvin.
    :param neutrons_produced_per_second: int or float - neutron flux.
    :param threshold: int or float - threshold for category.
    :return: str - one of ('LOW', 'NORMAL', 'DANGER').

    1. 'LOW' -> `temperature * neutrons per second` < 90% of `threshold`
    2. 'NORMAL' -> `temperature * neutrons per second` +/- 10% of `threshold`
    3. 'DANGER' -> `temperature * neutrons per second` is not in the above-stated ranges
    """
    # obtain a result  
    result = temperature * neutrons_produced_per_second 
    
    # define a lower bound 
    low = result < threshold * get_percent(90)

    # define a range 
    plus, minus = get_plus_or_minus(threshold, get_percent(10)) 

    # define normal based on range 
    upper = result < plus
    lower = result > minus
    normal = lower and upper
    
    # define a dangerous range
    danger = result > threshold * get_percent(90)

    failure = ""

    if low :
        failure = "LOW"
    elif normal:
        failure = "NORMAL"
    elif danger:
        failure = "DANGER"

    return failure

def get_percent(amount):
    """
    Get percent is a helper function

    :param amount: int or float - value to be made into a percent
    :return: float - amount is returned as a float , eg percentage of 100   
    """
    return amount / 100

def get_plus_or_minus(base, percent):
    """
    Get plus or minus is a helper 

    :param base: int - the base for which to get plus or minus
    :param percent: float - the percentage to asses the base against
    :return: float , float - the plus and minus of the base per its percentage
    """

    return base + (base * percent) , base - (base * percent)
