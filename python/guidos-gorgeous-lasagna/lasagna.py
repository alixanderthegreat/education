"""
Each Lasagna takes 40 minutes
"""
EXPECTED_BAKE_TIME = 40

"""
Each Layer takes 2 minutes
"""
PREPARATION_TIME = 2

def bake_time_remaining(elapsed_bake_time):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """

    return EXPECTED_BAKE_TIME - elapsed_bake_time

def preparation_time_in_minutes(layers):
    """
    Calculate the time necessary for size of lasagna

    :param layers: int - number of lasagna layers 
    :return: int - time needed to have prepared lasagna
    """
    return PREPARATION_TIME * layers

def elapsed_time_in_minutes(layers, time):
    """
    Calculate the total time used for baking lasagna

    :param layers: int - the number of layers of lasagna
    :param time: int - the amount of time that has already baked
    :return: int - the time that has elapsed
    """
    return (layers * PREPARATION_TIME) + time
