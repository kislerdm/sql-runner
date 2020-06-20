# SQL Runner

A routine to run SQL query/ies in RedShift.

## Objective 

To build a generic light-weight application to *synchronously* run in a k8s pod to execute submitted queries in Redshift.

## Motivation

Data preparation is often ongoing in a redshift cluster across many organisations. Since a couple years, kubernetes/k8s is being leveraged as the foundation for organisation-wide data platform. 

*Python* is a great programming language, and de-facto the default technology for most of data engineering teams. The services they made 

## Steps

1. Connect to s3
2. Conenct to RedShift
3. Read query from s3 bucket
4. Impute query paramters 
5. Split query to queries
6. Loop over queries list and execute every sub-query
7. Close connections
