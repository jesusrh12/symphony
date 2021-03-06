#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field as _field
from ...config import custom_scalars, datetime
from gql_client.runtime.variables import encode_variables
from gql import gql, Client
from gql.transport.exceptions import TransportQueryError
from functools import partial
from numbers import Number
from typing import Any, AsyncGenerator, Dict, List, Generator, Optional
from time import perf_counter
from dataclasses_json import DataClassJsonMixin, config

from .property import PropertyFragment, QUERY as PropertyFragmentQuery

# fmt: off
QUERY: List[str] = PropertyFragmentQuery + ["""
fragment LocationFragment on Location {
  id
  name
  latitude
  longitude
  externalId
  locationType {
    name
  }
  properties {
    ...PropertyFragment
  }
}

"""]

@dataclass(frozen=True)
class LocationFragment(DataClassJsonMixin):
    @dataclass(frozen=True)
    class LocationType(DataClassJsonMixin):
        name: str

    @dataclass(frozen=True)
    class Property(PropertyFragment):
        pass

    id: str
    name: str
    latitude: Number
    longitude: Number
    externalId: Optional[str]
    locationType: LocationType
    properties: List[Property]
