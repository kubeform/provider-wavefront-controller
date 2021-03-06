//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dashboard) DeepCopyInto(out *Dashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dashboard.
func (in *Dashboard) DeepCopy() *Dashboard {
	if in == nil {
		return nil
	}
	out := new(Dashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardList) DeepCopyInto(out *DashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardList.
func (in *DashboardList) DeepCopy() *DashboardList {
	if in == nil {
		return nil
	}
	out := new(DashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpec) DeepCopyInto(out *DashboardSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DashboardSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpec.
func (in *DashboardSpec) DeepCopy() *DashboardSpec {
	if in == nil {
		return nil
	}
	out := new(DashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecParameterDetails) DeepCopyInto(out *DashboardSpecParameterDetails) {
	*out = *in
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(string)
		**out = **in
	}
	if in.DynamicFieldType != nil {
		in, out := &in.DynamicFieldType, &out.DynamicFieldType
		*out = new(string)
		**out = **in
	}
	if in.HideFromView != nil {
		in, out := &in.HideFromView, &out.HideFromView
		*out = new(bool)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ParameterType != nil {
		in, out := &in.ParameterType, &out.ParameterType
		*out = new(string)
		**out = **in
	}
	if in.QueryValue != nil {
		in, out := &in.QueryValue, &out.QueryValue
		*out = new(string)
		**out = **in
	}
	if in.TagKey != nil {
		in, out := &in.TagKey, &out.TagKey
		*out = new(string)
		**out = **in
	}
	if in.ValuesToReadableStrings != nil {
		in, out := &in.ValuesToReadableStrings, &out.ValuesToReadableStrings
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecParameterDetails.
func (in *DashboardSpecParameterDetails) DeepCopy() *DashboardSpecParameterDetails {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecParameterDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecResource) DeepCopyInto(out *DashboardSpecResource) {
	*out = *in
	if in.CanModify != nil {
		in, out := &in.CanModify, &out.CanModify
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CanView != nil {
		in, out := &in.CanView, &out.CanView
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayQueryParameters != nil {
		in, out := &in.DisplayQueryParameters, &out.DisplayQueryParameters
		*out = new(bool)
		**out = **in
	}
	if in.DisplaySectionTableOfContents != nil {
		in, out := &in.DisplaySectionTableOfContents, &out.DisplaySectionTableOfContents
		*out = new(bool)
		**out = **in
	}
	if in.EventFilterType != nil {
		in, out := &in.EventFilterType, &out.EventFilterType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ParameterDetails != nil {
		in, out := &in.ParameterDetails, &out.ParameterDetails
		*out = make([]DashboardSpecParameterDetails, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Section != nil {
		in, out := &in.Section, &out.Section
		*out = make([]DashboardSpecSection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecResource.
func (in *DashboardSpecResource) DeepCopy() *DashboardSpecResource {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecSection) DeepCopyInto(out *DashboardSpecSection) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Row != nil {
		in, out := &in.Row, &out.Row
		*out = make([]DashboardSpecSectionRow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecSection.
func (in *DashboardSpecSection) DeepCopy() *DashboardSpecSection {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecSectionRow) DeepCopyInto(out *DashboardSpecSectionRow) {
	*out = *in
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = make([]DashboardSpecSectionRowChart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecSectionRow.
func (in *DashboardSpecSectionRow) DeepCopy() *DashboardSpecSectionRow {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecSectionRow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecSectionRowChart) DeepCopyInto(out *DashboardSpecSectionRowChart) {
	*out = *in
	if in.Base != nil {
		in, out := &in.Base, &out.Base
		*out = new(int64)
		**out = **in
	}
	if in.ChartAttribute != nil {
		in, out := &in.ChartAttribute, &out.ChartAttribute
		*out = new(string)
		**out = **in
	}
	if in.ChartSetting != nil {
		in, out := &in.ChartSetting, &out.ChartSetting
		*out = new(DashboardSpecSectionRowChartChartSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]DashboardSpecSectionRowChartSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Summarization != nil {
		in, out := &in.Summarization, &out.Summarization
		*out = new(string)
		**out = **in
	}
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecSectionRowChart.
func (in *DashboardSpecSectionRowChart) DeepCopy() *DashboardSpecSectionRowChart {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecSectionRowChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecSectionRowChartChartSetting) DeepCopyInto(out *DashboardSpecSectionRowChartChartSetting) {
	*out = *in
	if in.AutoColumnTags != nil {
		in, out := &in.AutoColumnTags, &out.AutoColumnTags
		*out = new(bool)
		**out = **in
	}
	if in.ColumnTags != nil {
		in, out := &in.ColumnTags, &out.ColumnTags
		*out = new(string)
		**out = **in
	}
	if in.CustomTags != nil {
		in, out := &in.CustomTags, &out.CustomTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExpectedDataSpacing != nil {
		in, out := &in.ExpectedDataSpacing, &out.ExpectedDataSpacing
		*out = new(int64)
		**out = **in
	}
	if in.FixedLegendDisplayStats != nil {
		in, out := &in.FixedLegendDisplayStats, &out.FixedLegendDisplayStats
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FixedLegendEnabled != nil {
		in, out := &in.FixedLegendEnabled, &out.FixedLegendEnabled
		*out = new(bool)
		**out = **in
	}
	if in.FixedLegendFilterField != nil {
		in, out := &in.FixedLegendFilterField, &out.FixedLegendFilterField
		*out = new(string)
		**out = **in
	}
	if in.FixedLegendFilterLimit != nil {
		in, out := &in.FixedLegendFilterLimit, &out.FixedLegendFilterLimit
		*out = new(int64)
		**out = **in
	}
	if in.FixedLegendFilterSort != nil {
		in, out := &in.FixedLegendFilterSort, &out.FixedLegendFilterSort
		*out = new(string)
		**out = **in
	}
	if in.FixedLegendHideLabel != nil {
		in, out := &in.FixedLegendHideLabel, &out.FixedLegendHideLabel
		*out = new(bool)
		**out = **in
	}
	if in.FixedLegendPosition != nil {
		in, out := &in.FixedLegendPosition, &out.FixedLegendPosition
		*out = new(string)
		**out = **in
	}
	if in.FixedLegendUseRawStats != nil {
		in, out := &in.FixedLegendUseRawStats, &out.FixedLegendUseRawStats
		*out = new(bool)
		**out = **in
	}
	if in.GroupBySource != nil {
		in, out := &in.GroupBySource, &out.GroupBySource
		*out = new(bool)
		**out = **in
	}
	if in.InvertDynamicLegendHoverControl != nil {
		in, out := &in.InvertDynamicLegendHoverControl, &out.InvertDynamicLegendHoverControl
		*out = new(bool)
		**out = **in
	}
	if in.LineType != nil {
		in, out := &in.LineType, &out.LineType
		*out = new(string)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.NumTags != nil {
		in, out := &in.NumTags, &out.NumTags
		*out = new(int64)
		**out = **in
	}
	if in.PlainMarkdownContent != nil {
		in, out := &in.PlainMarkdownContent, &out.PlainMarkdownContent
		*out = new(string)
		**out = **in
	}
	if in.ShowHosts != nil {
		in, out := &in.ShowHosts, &out.ShowHosts
		*out = new(bool)
		**out = **in
	}
	if in.ShowLabels != nil {
		in, out := &in.ShowLabels, &out.ShowLabels
		*out = new(bool)
		**out = **in
	}
	if in.ShowRawValues != nil {
		in, out := &in.ShowRawValues, &out.ShowRawValues
		*out = new(bool)
		**out = **in
	}
	if in.SortValuesDescending != nil {
		in, out := &in.SortValuesDescending, &out.SortValuesDescending
		*out = new(bool)
		**out = **in
	}
	if in.SparklineDecimalPrecision != nil {
		in, out := &in.SparklineDecimalPrecision, &out.SparklineDecimalPrecision
		*out = new(int64)
		**out = **in
	}
	if in.SparklineDisplayColor != nil {
		in, out := &in.SparklineDisplayColor, &out.SparklineDisplayColor
		*out = new(string)
		**out = **in
	}
	if in.SparklineDisplayFontSize != nil {
		in, out := &in.SparklineDisplayFontSize, &out.SparklineDisplayFontSize
		*out = new(string)
		**out = **in
	}
	if in.SparklineDisplayHorizontalPosition != nil {
		in, out := &in.SparklineDisplayHorizontalPosition, &out.SparklineDisplayHorizontalPosition
		*out = new(string)
		**out = **in
	}
	if in.SparklineDisplayPostfix != nil {
		in, out := &in.SparklineDisplayPostfix, &out.SparklineDisplayPostfix
		*out = new(string)
		**out = **in
	}
	if in.SparklineDisplayPrefix != nil {
		in, out := &in.SparklineDisplayPrefix, &out.SparklineDisplayPrefix
		*out = new(string)
		**out = **in
	}
	if in.SparklineDisplayValueType != nil {
		in, out := &in.SparklineDisplayValueType, &out.SparklineDisplayValueType
		*out = new(string)
		**out = **in
	}
	if in.SparklineDisplayVerticalPosition != nil {
		in, out := &in.SparklineDisplayVerticalPosition, &out.SparklineDisplayVerticalPosition
		*out = new(string)
		**out = **in
	}
	if in.SparklineFillColor != nil {
		in, out := &in.SparklineFillColor, &out.SparklineFillColor
		*out = new(string)
		**out = **in
	}
	if in.SparklineLineColor != nil {
		in, out := &in.SparklineLineColor, &out.SparklineLineColor
		*out = new(string)
		**out = **in
	}
	if in.SparklineSize != nil {
		in, out := &in.SparklineSize, &out.SparklineSize
		*out = new(string)
		**out = **in
	}
	if in.SparklineValueColorMapApplyTo != nil {
		in, out := &in.SparklineValueColorMapApplyTo, &out.SparklineValueColorMapApplyTo
		*out = new(string)
		**out = **in
	}
	if in.SparklineValueColorMapColors != nil {
		in, out := &in.SparklineValueColorMapColors, &out.SparklineValueColorMapColors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SparklineValueColorMapValues != nil {
		in, out := &in.SparklineValueColorMapValues, &out.SparklineValueColorMapValues
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.SparklineValueColorMapValuesV2 != nil {
		in, out := &in.SparklineValueColorMapValuesV2, &out.SparklineValueColorMapValuesV2
		*out = make([]float64, len(*in))
		copy(*out, *in)
	}
	if in.SparklineValueTextMapText != nil {
		in, out := &in.SparklineValueTextMapText, &out.SparklineValueTextMapText
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SparklineValueTextMapThresholds != nil {
		in, out := &in.SparklineValueTextMapThresholds, &out.SparklineValueTextMapThresholds
		*out = make([]float64, len(*in))
		copy(*out, *in)
	}
	if in.StackType != nil {
		in, out := &in.StackType, &out.StackType
		*out = new(string)
		**out = **in
	}
	if in.TagMode != nil {
		in, out := &in.TagMode, &out.TagMode
		*out = new(string)
		**out = **in
	}
	if in.TimeBasedColoring != nil {
		in, out := &in.TimeBasedColoring, &out.TimeBasedColoring
		*out = new(bool)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.WindowSize != nil {
		in, out := &in.WindowSize, &out.WindowSize
		*out = new(int64)
		**out = **in
	}
	if in.Windowing != nil {
		in, out := &in.Windowing, &out.Windowing
		*out = new(string)
		**out = **in
	}
	if in.Xmax != nil {
		in, out := &in.Xmax, &out.Xmax
		*out = new(float64)
		**out = **in
	}
	if in.Xmin != nil {
		in, out := &in.Xmin, &out.Xmin
		*out = new(float64)
		**out = **in
	}
	if in.Y0ScaleSiBy1024 != nil {
		in, out := &in.Y0ScaleSiBy1024, &out.Y0ScaleSiBy1024
		*out = new(bool)
		**out = **in
	}
	if in.Y0UnitAutoscaling != nil {
		in, out := &in.Y0UnitAutoscaling, &out.Y0UnitAutoscaling
		*out = new(bool)
		**out = **in
	}
	if in.Y1ScaleSiBy1024 != nil {
		in, out := &in.Y1ScaleSiBy1024, &out.Y1ScaleSiBy1024
		*out = new(bool)
		**out = **in
	}
	if in.Y1UnitAutoscaling != nil {
		in, out := &in.Y1UnitAutoscaling, &out.Y1UnitAutoscaling
		*out = new(bool)
		**out = **in
	}
	if in.Y1Units != nil {
		in, out := &in.Y1Units, &out.Y1Units
		*out = new(string)
		**out = **in
	}
	if in.Y1max != nil {
		in, out := &in.Y1max, &out.Y1max
		*out = new(float64)
		**out = **in
	}
	if in.Y1min != nil {
		in, out := &in.Y1min, &out.Y1min
		*out = new(float64)
		**out = **in
	}
	if in.Ymax != nil {
		in, out := &in.Ymax, &out.Ymax
		*out = new(float64)
		**out = **in
	}
	if in.Ymin != nil {
		in, out := &in.Ymin, &out.Ymin
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecSectionRowChartChartSetting.
func (in *DashboardSpecSectionRowChartChartSetting) DeepCopy() *DashboardSpecSectionRowChartChartSetting {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecSectionRowChartChartSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecSectionRowChartSource) DeepCopyInto(out *DashboardSpecSectionRowChartSource) {
	*out = *in
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.QueryBuilderEnabled != nil {
		in, out := &in.QueryBuilderEnabled, &out.QueryBuilderEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ScatterPlotSource != nil {
		in, out := &in.ScatterPlotSource, &out.ScatterPlotSource
		*out = new(string)
		**out = **in
	}
	if in.SourceDescription != nil {
		in, out := &in.SourceDescription, &out.SourceDescription
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecSectionRowChartSource.
func (in *DashboardSpecSectionRowChartSource) DeepCopy() *DashboardSpecSectionRowChartSource {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecSectionRowChartSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardStatus) DeepCopyInto(out *DashboardStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardStatus.
func (in *DashboardStatus) DeepCopy() *DashboardStatus {
	if in == nil {
		return nil
	}
	out := new(DashboardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Json) DeepCopyInto(out *Json) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Json.
func (in *Json) DeepCopy() *Json {
	if in == nil {
		return nil
	}
	out := new(Json)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Json) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonList) DeepCopyInto(out *JsonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Json, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonList.
func (in *JsonList) DeepCopy() *JsonList {
	if in == nil {
		return nil
	}
	out := new(JsonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JsonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonSpec) DeepCopyInto(out *JsonSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(JsonSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonSpec.
func (in *JsonSpec) DeepCopy() *JsonSpec {
	if in == nil {
		return nil
	}
	out := new(JsonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonSpecResource) DeepCopyInto(out *JsonSpecResource) {
	*out = *in
	if in.DashboardJSON != nil {
		in, out := &in.DashboardJSON, &out.DashboardJSON
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonSpecResource.
func (in *JsonSpecResource) DeepCopy() *JsonSpecResource {
	if in == nil {
		return nil
	}
	out := new(JsonSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsonStatus) DeepCopyInto(out *JsonStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsonStatus.
func (in *JsonStatus) DeepCopy() *JsonStatus {
	if in == nil {
		return nil
	}
	out := new(JsonStatus)
	in.DeepCopyInto(out)
	return out
}
