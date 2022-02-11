#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field as _field
from functools import partial
from ...config import custom_scalars, datetime
from numbers import Number
from typing import Any, AsyncGenerator, Dict, List, Generator, Optional

from dataclasses_json import DataClassJsonMixin, config

from gql_client.runtime.enum_utils import enum_field_metadata
from ..enum.project_priority import ProjectPriority

from ..input.property_input import PropertyInput


@dataclass(frozen=True)
class EditProjectInput(DataClassJsonMixin):
    id: str
    name: str
    type: str
    properties: List[PropertyInput]
    description: Optional[str] = None
    priority: Optional[ProjectPriority] = None
    creatorId: Optional[str] = None
    location: Optional[str] = None
